package strategy

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TransactionBuilderFunc is a function that returns an ethereum transaction
type TransactionBuilderFunc func() (*types.Transaction, error)

type Arguments struct {
	Backend         bind.ContractBackend
	Key             *ecdsa.PrivateKey
	ChainID         *big.Int
	Strategy        *strategy.Strategy
	Safe            common.Address
	Aggregator      aggregators.Aggregator
	StrategyAddress common.Address
	StableCoin      common.Address
	DepositToken    common.Address
	NoSellTokens    []common.Address
	Pool            *big.Int
}

// NewHarvestFunc returns a function that builds a harvest transaction with swap parameters
func NewHarvestFunc(args Arguments) TransactionBuilderFunc {
	return func() (*types.Transaction, error) {
		return executeAction(args, args.Strategy.Harvest)
	}
}

// NewCompoundFunc returns a function that builds a compound transaction with swap parameters
func NewCompoundFunc(args Arguments) TransactionBuilderFunc {
	return func() (*types.Transaction, error) {

		// We change the `StableCoin` to depositToken.
		// This means our code will ignore any harvests that are the depositToken for selling and try to swap all others
		// into the depositToken.
		args.StableCoin = args.DepositToken

		return executeAction(args, args.Strategy.Compound)
	}
}

type actionFunc func(*bind.TransactOpts, *big.Int, common.Address, []byte) (*types.Transaction, error)

type SwapParams struct {
	Router       common.Address
	Spender      common.Address
	Input        common.Address
	AmountIn     *big.Int
	Output       common.Address
	AmountOutMin *big.Int
	Data         []byte
}

func executeAction(args Arguments, action actionFunc) (*types.Transaction, error) {
	harvests, err := SimulateClaim(args.Backend, args.StrategyAddress, args.Pool, args.Safe)
	if err != nil {
		return nil, err
	}

	harvests = filter(harvests, append(args.NoSellTokens, args.StableCoin))

	var swaps []*aggregators.Swap
	for _, harvest := range harvests {
		swap, err := args.Aggregator.GetSwap(aggregators.SwapArgs{
			ChainID:     args.ChainID.Uint64(),
			From:        harvest.Token,
			To:          args.StableCoin,
			FromAddress: args.Safe,
			Amount:      harvest.Amount,
			Slippage:    big.NewFloat(1.0),
			Opts: aggregators.OneInchSwapOpts{
				DisableEstimate: true,
			},
		})
		if err != nil {
			return nil, err
		}

		swaps = append(swaps, swap)
	}

	swapsParams := transformSwap(swaps)
	swapsParamsEncoded, err := EncodeSwapParams(swapsParams)
	if err != nil {
		return nil, err
	}

	opts, err := newKeyedTransactorWithChainIDAndNoSend(args.Key, args.ChainID)
	if err != nil {
		return nil, err
	}

	tx, err := action(opts, args.Pool, args.Safe, swapsParamsEncoded)
	if err != nil {
		return nil, fmt.Errorf("failed with err: %s safe: %s strategy: %s pool: %s data: %s", err, args.Safe.Hex(), args.StrategyAddress.Hex(), args.Pool.String(), common.Bytes2Hex(swapsParamsEncoded))
	}

	tx, err = increaseGasLimit(tx)
	if err != nil {
		return nil, err
	}

	signedTx, err := opts.Signer(opts.From, tx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func increaseGasLimit(tx *types.Transaction) (*types.Transaction, error) {
	// adds 20% to gas estimation.
	gas := (tx.Gas() * (1000 + 200)) / 1000

	switch tx.Type() {
	case types.LegacyTxType:
		return types.NewTx(&types.LegacyTx{
			Nonce:    tx.Nonce(),
			GasPrice: tx.GasPrice(),
			Gas:      gas,
			Value:    tx.Value(),
			Data:     tx.Data(),
			To:       tx.To(),
		}), nil
	case types.DynamicFeeTxType:
		return types.NewTx(&types.DynamicFeeTx{
			ChainID:   tx.ChainId(),
			Nonce:     tx.Nonce(),
			GasFeeCap: tx.GasFeeCap(),
			GasTipCap: tx.GasTipCap(),
			Gas:       gas,
			Value:     tx.Value(),
			Data:      tx.Data(),
			To:        tx.To(),
		}), nil
	}

	return nil, fmt.Errorf("unknown transaction type %d", tx.Type())
}

func newKeyedTransactorWithChainIDAndNoSend(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}

	opts.NoSend = true

	return opts, nil
}

func transformSwap(vs []*aggregators.Swap) (ret []SwapParams) {
	for _, v := range vs {
		slippedAmount := utils.SlippageMillis(v.ToTokenAmount, 10)

		ret = append(ret, SwapParams{
			Router:       v.Tx.To,
			Spender:      v.Tx.To,
			Input:        v.FromToken,
			AmountIn:     v.FromTokenAmount,
			Output:       v.ToToken,
			AmountOutMin: slippedAmount,
			Data:         v.Tx.Data,
		})
	}

	return ret
}

func EncodeSwapParams(vs []SwapParams) ([]byte, error) {
	swapsType, err := abi.NewType("tuple[]", "SwapParams[]", []abi.ArgumentMarshaling{
		{Name: "router", Type: "address"},
		{Name: "spender", Type: "address"},
		{Name: "input", Type: "address"},
		{Name: "amountIn", Type: "uint256"},
		{Name: "output", Type: "address"},
		{Name: "amountOutMin", Type: "uint256"},
		{Name: "data", Type: "bytes"},
	})

	if err != nil {
		return nil, err
	}

	swapsArgs := abi.Arguments{
		{Name: "swaps", Type: swapsType},
	}

	return swapsArgs.Pack(vs)
}

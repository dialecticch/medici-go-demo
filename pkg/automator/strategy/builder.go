package strategy

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/dialecticch/medici-go/pkg/automator/fees"
	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Builder struct {
	key     *ecdsa.PrivateKey
	chainID *big.Int

	backend bind.ContractBackend

	stablecoinDecimals int
	stablecoin         common.Address

	wrappedToken common.Address
	aggregator   aggregators.Aggregator
	noSellTokens []common.Address

	calculator fees.Calculator
}

func NewBuilder(
	key *ecdsa.PrivateKey,
	chainID *big.Int,
	backend bind.ContractBackend,
	stablecoinDecimals int,
	stablecoin common.Address,
	wrappedToken common.Address,
	calculator fees.Calculator,
	aggregator aggregators.Aggregator,
	noSellTokens []common.Address,
) *Builder {
	return &Builder{
		key:     key,
		chainID: chainID,

		backend: backend,

		stablecoinDecimals: stablecoinDecimals,
		stablecoin:         stablecoin,
		wrappedToken:       wrappedToken,
		aggregator:         aggregator,
		noSellTokens:       noSellTokens,

		calculator: calculator,
	}
}

func (b *Builder) Build(config *types.StrategyConfig, block *big.Int) (*Strategy, error) {

	s := &Strategy{}
	s.addr = config.Strategy.Address
	s.safe = config.Safe
	s.pool = config.Pool

	strat, err := strategy.NewStrategy(config.Strategy.Address, b.backend)
	if err != nil {
		return nil, err
	}

	args := Arguments{
		Backend:         b.backend,
		Key:             b.key,
		ChainID:         b.chainID,
		Strategy:        strat,
		Safe:            config.Safe,
		Aggregator:      b.aggregator,
		StrategyAddress: config.Strategy.Address,
		StableCoin:      b.stablecoin,
		DepositToken:    config.DepositToken,
		NoSellTokens:    b.noSellTokens,
		Pool:            config.Pool,
	}

	tb, err := b.buildTransactionBuilder(config.Action, args)
	if err != nil {
		return nil, err
	}

	s.transactionBuilder = tb

	t, err := b.buildThreshold(config, block)
	if err != nil {
		return nil, err
	}

	s.threshold = t

	return s, nil
}

func (b *Builder) buildTransactionBuilder(action string, args Arguments) (TransactionBuilderFunc, error) {
	switch action {
	case "harvest":
		return NewHarvestFunc(args), nil
	case "compound":
		return NewCompoundFunc(args), nil
	default:
		return nil, fmt.Errorf("no transaction builder for %s", action)
	}
}

func (b *Builder) buildThreshold(config *types.StrategyConfig, block *big.Int) (ThresholdFunc, error) {
	switch config.ThresholdType {
	case "gas_percentage":
		return NewGasThresholdFunc(
			GasThresholdArguments{
				Backend:         b.backend,
				StrategyAddress: config.Strategy.Address,
				ChainID:         b.chainID.Uint64(),
				Aggregator:      b.aggregator,
				Calculator:      b.calculator,
				Decimals:        b.stablecoinDecimals,
				Threshold:       config.Threshold,
				Safe:            config.Safe,
				StableCoin:      b.stablecoin,
				Pool:            config.Pool,
			},
		), nil
	case "elapsed_blocks":
		return NewElapsedThresholdFunc(b.backend, config.Threshold, block, config.LastHarvested), nil
	default:
		return nil, fmt.Errorf("no threshold for %s", config.ThresholdType)
	}
}

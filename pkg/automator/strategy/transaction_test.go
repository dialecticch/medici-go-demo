package strategy_test

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/dialecticch/medici-go/pkg/automator/strategy"
	"github.com/dialecticch/medici-go/pkg/automator/testdata"
	contract "github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/testutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
)

var (
	routerAddress = common.HexToAddress("0x1111111254fb6c44bAC0beD2854e76F90643097d") // 1inch v4 router
)

func TestHarvestFunc(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	ctrl := gomock.NewController(t)

	aggregator := testdata.NewMockAggregator(ctrl)

	b := testutils.SetupBackend(privateKey)
	chainID := big.NewInt(1337)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	strategyAddress, _, _, err := testutils.DeployMockStrategy(transactOpts, b)
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	b.Commit()

	strat, err := contract.NewStrategy(strategyAddress, b)
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	// Expect a call to the aggregator, getting quote for `wrapped` of amount 10 (see the mock strategy)
	aggregator.EXPECT().GetSwap(aggregators.SwapArgs{
		ChainID:     1337,
		From:        common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		To:          common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		FromAddress: common.HexToAddress("0x00000000000000000000000000000000000054FE"),
		Amount:      big.NewInt(10),
		Slippage:    big.NewFloat(1.0),
		Opts: aggregators.OneInchSwapOpts{
			DisableEstimate: true,
		},
	}).Return(
		&aggregators.Swap{
			FromToken:       common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			ToToken:         common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			ToTokenAmount:   big.NewInt(100000000),
			FromTokenAmount: big.NewInt(10),
			Tx: &aggregators.AggregatorTransaction{
				From:     common.HexToAddress("0x00000000000000000000000000000000000054FE"),
				To:       routerAddress,
				Data:     []byte{0},
				Value:    big.NewInt(0),
				GasPrice: big.NewInt(0),
				Gas:      150_000,
			},
		},
		nil,
	)

	signedTx, err := strategy.NewHarvestFunc(strategy.Arguments{
		Backend:         b,
		Key:             privateKey,
		ChainID:         chainID,
		Strategy:        strat,
		Safe:            common.HexToAddress("0x00000000000000000000000000000000000054FE"),
		Aggregator:      aggregator,
		StrategyAddress: strategyAddress,
		StableCoin:      common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		NoSellTokens:    []common.Address{},
		Pool:            common.Big1,
	})()

	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	expected, _ := strategy.EncodeSwapParams([]strategy.SwapParams{
		{
			Router:       routerAddress,
			Spender:      routerAddress,
			Input:        common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			AmountIn:     big.NewInt(10),
			Output:       common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			AmountOutMin: big.NewInt(99000000),
			Data:         []byte{0},
		},
	})

	data := signedTx.Data()
	if !bytes.Equal(data[len(data)-len(expected):], expected) {
		t.Fatal("Tx contains wrong swaps. Expected ", expected, " got ", data[len(data)-len(expected):])
	}
}

func TestCompoundFunc(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	ctrl := gomock.NewController(t)

	aggregator := testdata.NewMockAggregator(ctrl)

	b := testutils.SetupBackend(privateKey)
	chainID := big.NewInt(1337)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	strategyAddress, _, _, err := testutils.DeployMockStrategy(transactOpts, b)
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	b.Commit()

	strat, err := contract.NewStrategy(strategyAddress, b)
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	// Expect a call to the aggregator, getting quote for `wrapped` of amount 10 (see the mock strategy)
	aggregator.EXPECT().GetSwap(aggregators.SwapArgs{
		ChainID:     1337,
		From:        common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		To:          common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		FromAddress: common.HexToAddress("0x00000000000000000000000000000000000054FE"),
		Amount:      big.NewInt(10),
		Slippage:    big.NewFloat(1.0),
		Opts: aggregators.OneInchSwapOpts{
			DisableEstimate: true,
		},
	}).Return(
		&aggregators.Swap{
			FromToken:       common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			ToToken:         common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			ToTokenAmount:   big.NewInt(100000000),
			FromTokenAmount: big.NewInt(10),
			Tx: &aggregators.AggregatorTransaction{
				From:     common.HexToAddress("0x00000000000000000000000000000000000054FE"),
				To:       routerAddress,
				Data:     []byte{0},
				Value:    big.NewInt(0),
				GasPrice: big.NewInt(0),
				Gas:      150_000,
			},
		},
		nil,
	)

	signedTx, err := strategy.NewHarvestFunc(strategy.Arguments{
		Backend:         b,
		Key:             privateKey,
		ChainID:         chainID,
		Strategy:        strat,
		Safe:            common.HexToAddress("0x00000000000000000000000000000000000054FE"),
		Aggregator:      aggregator,
		StrategyAddress: strategyAddress,
		StableCoin:      common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		NoSellTokens:    []common.Address{},
		Pool:            common.Big1,
	})()

	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	expected, _ := strategy.EncodeSwapParams([]strategy.SwapParams{
		{
			Router:       routerAddress,
			Spender:      routerAddress,
			Input:        common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			AmountIn:     big.NewInt(10),
			Output:       common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			AmountOutMin: big.NewInt(99000000),
			Data:         []byte{0},
		},
	})

	data := signedTx.Data()
	if !bytes.Equal(data[len(data)-len(expected):], expected) {
		t.Fatal("Tx contains wrong swaps. Expected ", expected, " got ", data[len(data)-len(expected):])
	}
}

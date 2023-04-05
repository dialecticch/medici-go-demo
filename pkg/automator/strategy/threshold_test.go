package strategy_test

import (
	"math/big"
	"testing"

	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/dialecticch/medici-go/pkg/automator/fees"
	"github.com/dialecticch/medici-go/pkg/automator/strategy"
	"github.com/dialecticch/medici-go/pkg/automator/testdata"
	"github.com/dialecticch/medici-go/pkg/testutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
)

var (
	safeAddress = common.HexToAddress("0x00000000000000000000000000000000000054FE")
	wrapped     = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2") // wrapped native token
	stableCoin  = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // mainnet usdc address
)

func TestGasThreshold_HappyPath(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)

	oracle := testdata.NewMockPriceOracle(ctrl)
	aggregator := testdata.NewMockAggregator(ctrl)

	b := testutils.SetupBackend(privateKey)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	strategyAddress, _, _, err := testutils.DeployMockStrategy(transactOpts, b)

	b.Commit()

	if err != nil {
		t.Fatal(err)
	}

	// Expect a call to the aggregator, getting quote for `wrapped` of amount 10 (see the mock strategy)
	aggregator.EXPECT().GetQuote(aggregators.QuoteArgs{
		ChainID: 1337,
		From:    wrapped,
		To:      stableCoin,
		Amount:  big.NewInt(10),
		Opts:    aggregators.OneInchQuoteOpts{},
	}).Return(
		&aggregators.Quote{
			FromToken:       wrapped,
			ToToken:         stableCoin,
			ToTokenAmount:   big.NewInt(100000000),
			FromTokenAmount: big.NewInt(10),
			EstimatedGas:    0,
		},
		nil,
	)

	// Expect calls to the oracle (we reached the evaluation of the harvest)
	oracle.EXPECT().Decimals().Return(6)
	oracle.EXPECT().Price().Return(big.NewInt(1))

	result, err := strategy.NewGasThresholdFunc(strategy.GasThresholdArguments{
		Backend:         b,
		StrategyAddress: strategyAddress,
		ChainID:         1337,
		Aggregator:      aggregator,
		Calculator:      fees.NewEthereumFeeCalculator(oracle),
		Decimals:        6,
		Threshold:       0.15,
		Safe:            safeAddress,
		StableCoin:      stableCoin,
		NoSellTokens:    []common.Address{},
		Pool:            common.Big1,
	},
	)(types.NewTransaction(0, strategyAddress, big.NewInt(0), 10, big.NewInt(10_000_00), []byte{}))

	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	if !result {
		t.Fatal("Expected true.")
	}
}

func TestGasThreshold_EarlyReturn(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)

	oracle := testdata.NewMockPriceOracle(ctrl)
	aggregator := testdata.NewMockAggregator(ctrl)

	b := testutils.SetupBackend(privateKey)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	strategyAddress, _, _, err := testutils.DeployMockStrategy(transactOpts, b)

	b.Commit()

	if err != nil {
		t.Fatal(err)
	}

	// Don't expect anything from the mocks.
	// We're not reaching the evaluation of the harvest as we're not doing swaps (`wrapped` and `stablecoin` are filtered out)
	result, err := strategy.NewGasThresholdFunc(strategy.GasThresholdArguments{
		Backend:         b,
		StrategyAddress: strategyAddress,
		ChainID:         1337,
		Aggregator:      aggregator,
		Calculator:      fees.NewEthereumFeeCalculator(oracle),
		Decimals:        5,
		Threshold:       0.15,
		Safe:            safeAddress,
		StableCoin:      stableCoin,
		NoSellTokens:    []common.Address{wrapped},
		Pool:            common.Big1,
	},
	)(&types.Transaction{})

	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	if result {
		t.Fatal("Expected false.")
	}
}

package automator_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/dialecticch/medici-go/pkg/automator/fees"
	"github.com/dialecticch/medici-go/pkg/automator/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"

	"github.com/dialecticch/medici-go/pkg/automator"
	"github.com/dialecticch/medici-go/pkg/automator/testdata"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/testutils"
)

var (
	stableCoinDecimals = 6
	chainID            = big.NewInt(1337)
	safeAddress        = common.HexToAddress("0x00000000000000000000000000000000000054FE")
	wrapped            = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2") // wrapped native token
	stableCoin         = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // mainnet usdc address
	routerAddress      = common.HexToAddress("0x1111111254fb6c44bAC0beD2854e76F90643097d") // 1inch v4 router
)

func TestAutomator_Run_With_GasThreshold(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	b := testutils.SetupBackend(privateKey)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	strategyAddress, _, _, err := testutils.DeployMockStrategy(transactOpts, b)
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	b.Commit()

	// Setup mocks
	ctrl := gomock.NewController(t)

	ms := testdata.NewMockSender(ctrl)
	oracle := testdata.NewMockPriceOracle(ctrl)
	aggregator := testdata.NewMockAggregator(ctrl)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sink := make(chan *types.Header)

	mock.ExpectQuery("^SELECT (.+)").
		WithArgs(1337, true).
		WillReturnRows(sqlmock.NewRows([]string{"safe", "strategy", "name", "threshold", "action", "threshold_type", "last_harvested", "pool", "deposit_token"}).AddRow(safeAddress.String(), strategyAddress.String(), "foo", "0.15", "harvest", "gas_percentage", "0", "1", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"))

	repo := repositories.NewStrategyRepository(db, 1337)

	// Create the builder
	builder := strategy.NewBuilder(
		privateKey,
		chainID,
		b,
		stableCoinDecimals,
		stableCoin,
		wrapped,
		fees.NewEthereumFeeCalculator(oracle),
		aggregator,
		[]common.Address{},
	)

	// Create the automator
	a := automator.NewAutomator(
		sink,
		repo,
		builder,
		ms,
		privateKey,
		5,
		3,
		nil,
		b,
	)

	go func() {
		err := a.Run()
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(1 * time.Second)

	success := make(chan bool)

	// Expect calls to the aggregator for quote and swap
	aggregator.EXPECT().GetQuote(aggregators.QuoteArgs{
		ChainID: chainID.Uint64(),
		From:    wrapped,
		To:      stableCoin,
		Amount:  big.NewInt(10),
		Opts:    aggregators.OneInchQuoteOpts{},
	}).Return(&aggregators.Quote{
		FromToken:       wrapped,
		ToToken:         stableCoin,
		ToTokenAmount:   big.NewInt(100000000),
		FromTokenAmount: big.NewInt(10),
		EstimatedGas:    0,
	},
		nil,
	)

	aggregator.EXPECT().GetSwap(aggregators.SwapArgs{
		ChainID:     chainID.Uint64(),
		From:        wrapped,
		To:          stableCoin,
		FromAddress: safeAddress,
		Amount:      big.NewInt(10),
		Slippage:    big.NewFloat(1.0),
		Opts: aggregators.OneInchSwapOpts{
			DisableEstimate: true,
		},
	}).Return(&aggregators.Swap{
		FromToken:       common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		ToToken:         common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		ToTokenAmount:   big.NewInt(100000000),
		FromTokenAmount: big.NewInt(10),
		Tx: &aggregators.AggregatorTransaction{
			From:     safeAddress,
			To:       routerAddress,
			Data:     []byte{0},
			Value:    big.NewInt(0),
			GasPrice: big.NewInt(0),
			Gas:      150_000,
		},
	},
		nil,
	)

	oracle.EXPECT().Decimals().Return(6)
	oracle.EXPECT().Price().Return(big.NewInt(3000_000000))

	ms.EXPECT().Send(gomock.Any()).DoAndReturn(func(*types.Transaction) (*types.Receipt, error) {
		success <- true
		return &types.Receipt{Status: 1}, nil
	})

	sink <- new(types.Header)

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(15 * time.Second)
		timeout <- true
	}()

	select {
	case <-success:
		return
	case <-timeout:
		t.Fatal("Failed for timeout")
	}
}

func TestAutomator_Run_With_BlocksElapsed(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	b := testutils.SetupBackend(privateKey)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	strategyAddress, _, _, err := testutils.DeployMockStrategy(transactOpts, b)
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	b.Commit()

	// Setup mocks
	ctrl := gomock.NewController(t)

	ms := testdata.NewMockSender(ctrl)
	oracle := testdata.NewMockPriceOracle(ctrl)
	aggregator := testdata.NewMockAggregator(ctrl)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sink := make(chan *types.Header)

	mock.ExpectQuery("^SELECT (.+)").
		WithArgs(1337, true).
		WillReturnRows(sqlmock.NewRows([]string{"safe", "strategy", "name", "threshold", "action", "threshold_type", "last_harvested", "pool", "deposit_token"}).AddRow(safeAddress.String(), strategyAddress.String(), "foo", "5", "harvest", "elapsed_blocks", "0", "1", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"))
	mock.ExpectQuery("^SELECT (.+)").
		WithArgs(1337, true).
		WillReturnRows(sqlmock.NewRows([]string{"safe", "strategy", "name", "threshold", "action", "threshold_type", "last_harvested", "pool", "deposit_token"}).AddRow(safeAddress.String(), strategyAddress.String(), "foo", "5", "harvest", "elapsed_blocks", "0", "1", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"))

	repo := repositories.NewStrategyRepository(db, 1337)

	// Create the builder
	builder := strategy.NewBuilder(
		privateKey,
		chainID,
		b,
		stableCoinDecimals,
		stableCoin,
		wrapped,
		fees.NewEthereumFeeCalculator(oracle),
		aggregator,
		[]common.Address{},
	)

	a := automator.NewAutomator(
		sink,
		repo,
		builder,
		ms,
		privateKey,
		5,
		3,
		nil,
		b,
	)

	go func() {
		err := a.Run()
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(1 * time.Second)

	success := make(chan bool)

	// Expect calls to the aggregator for swap
	aggregator.EXPECT().GetSwap(aggregators.SwapArgs{
		ChainID:     chainID.Uint64(),
		From:        wrapped,
		To:          stableCoin,
		FromAddress: safeAddress,
		Amount:      big.NewInt(10),
		Slippage:    big.NewFloat(1.0),
		Opts: aggregators.OneInchSwapOpts{
			DisableEstimate: true,
		},
	}).Return(&aggregators.Swap{
		FromToken:       common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		ToToken:         common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		ToTokenAmount:   big.NewInt(100000000),
		FromTokenAmount: big.NewInt(10),
		Tx: &aggregators.AggregatorTransaction{
			From:     safeAddress,
			To:       routerAddress,
			Data:     []byte{0},
			Value:    big.NewInt(0),
			GasPrice: big.NewInt(0),
			Gas:      150_000,
		},
	},
		nil,
	).Times(2)

	ms.EXPECT().Send(gomock.Any()).DoAndReturn(func(*types.Transaction) (*types.Receipt, error) {
		success <- true
		return &types.Receipt{Status: 1}, nil
	})

	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond) // LOL BIG Number opts are too slow.
		sink <- &types.Header{
			Number: big.NewInt(int64(i)),
		}
	}

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(15 * time.Second)
		timeout <- true
	}()

	select {
	case <-success:
		return
	case <-timeout:
		t.Fatal("timed out")
	}

}

func TestAutomator_Run_LogsWhenTransactionFailed(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	b := testutils.SetupBackend(privateKey)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	strategyAddress, _, _, err := testutils.DeployMockStrategy(transactOpts, b)
	if err != nil {
		t.Fatal("Failed with error: ", err)
	}

	b.Commit()

	// Setup mocks
	ctrl := gomock.NewController(t)

	ms := testdata.NewMockSender(ctrl)
	logger := testdata.NewMockLogger(ctrl)
	oracle := testdata.NewMockPriceOracle(ctrl)
	aggregator := testdata.NewMockAggregator(ctrl)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sink := make(chan *types.Header)

	mock.ExpectQuery("^SELECT (.+)").
		WithArgs(1337, true).
		WillReturnRows(sqlmock.NewRows([]string{"safe", "strategy", "name", "threshold", "action", "threshold_type", "last_harvested", "pool", "deposit_token"}).AddRow(safeAddress.String(), strategyAddress.String(), "foo", "5", "harvest", "elapsed_blocks", "0", "1", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"))
	mock.ExpectQuery("^SELECT (.+)").
		WithArgs(1337, true).
		WillReturnRows(sqlmock.NewRows([]string{"safe", "strategy", "name", "threshold", "action", "threshold_type", "last_harvested", "pool", "deposit_token"}).AddRow(safeAddress.String(), strategyAddress.String(), "foo", "5", "harvest", "elapsed_blocks", "0", "1", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"))

	repo := repositories.NewStrategyRepository(db, 1337)

	// Create the builder
	builder := strategy.NewBuilder(
		privateKey,
		chainID,
		b,
		stableCoinDecimals,
		stableCoin,
		wrapped,
		fees.NewEthereumFeeCalculator(oracle),
		aggregator,
		[]common.Address{},
	)

	a := automator.NewAutomator(
		sink,
		repo,
		builder,
		ms,
		privateKey,
		5,
		3,
		logger,
		b,
	)

	go func() {
		err := a.Run()
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(1 * time.Second)

	success := make(chan bool)

	// Expect calls to the aggregator for swap
	aggregator.EXPECT().GetSwap(aggregators.SwapArgs{
		ChainID:     chainID.Uint64(),
		From:        wrapped,
		To:          stableCoin,
		FromAddress: safeAddress,
		Amount:      big.NewInt(10),
		Slippage:    big.NewFloat(1.0),
		Opts: aggregators.OneInchSwapOpts{
			DisableEstimate: true,
		},
	}).Return(&aggregators.Swap{
		FromToken:       common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		ToToken:         common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		ToTokenAmount:   big.NewInt(100000000),
		FromTokenAmount: big.NewInt(10),
		Tx: &aggregators.AggregatorTransaction{
			From:     safeAddress,
			To:       routerAddress,
			Data:     []byte{0},
			Value:    big.NewInt(0),
			GasPrice: big.NewInt(0),
			Gas:      150_000,
		},
	},
		nil,
	).Times(2)

	ms.EXPECT().Send(gomock.Any()).DoAndReturn(func(*types.Transaction) (*types.Receipt, error) {
		return &types.Receipt{Status: 0}, nil
	})

	logger.EXPECT().Failed(gomock.Any()).DoAndReturn(func(receipt *types.Receipt) error {
		success <- true
		return nil
	})

	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond) // LOL BIG Number opts are too slow.
		sink <- &types.Header{
			Number: big.NewInt(int64(i)),
		}
	}

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(15 * time.Second)
		timeout <- true
	}()

	select {
	case <-success:
		break
	case <-timeout:
		t.Fatal("timed out")
	}

}

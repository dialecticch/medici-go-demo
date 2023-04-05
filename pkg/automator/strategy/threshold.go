package strategy

import (
	"log"
	"math"
	"math/big"

	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/dialecticch/medici-go/pkg/automator/fees"
	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ThresholdFunc returns whether a given transaction has reached the threshold for execution
type ThresholdFunc func(*types.Transaction) (bool, error)

type GasThresholdArguments struct {
	Backend         bind.ContractBackend
	StrategyAddress common.Address
	ChainID         uint64
	Aggregator      aggregators.Aggregator
	Calculator      fees.Calculator
	Decimals        int
	Threshold       float64
	Safe            common.Address
	StableCoin      common.Address
	NoSellTokens    []common.Address
	Pool            *big.Int
}

func NewGasThresholdFunc(args GasThresholdArguments) ThresholdFunc {
	return func(tx *types.Transaction) (bool, error) {
		harvests, err := SimulateClaim(args.Backend, args.StrategyAddress, args.Pool, args.Safe)
		if err != nil {
			return false, err
		}

		harvests = filter(harvests, append(args.NoSellTokens, args.StableCoin))

		amount, err := getQuoteForHarvests(harvests, args.Aggregator, args.ChainID, args.StableCoin)
		if err != nil {
			return false, err
		}

		return calculateThresholdReached(amount, tx, args.Calculator, args.Decimals, args.Threshold), nil
	}
}

func NewElapsedThresholdFunc(_ bind.ContractBackend, threshold float64, block *big.Int, lastHarvested big.Int) ThresholdFunc {
	return func(*types.Transaction) (bool, error) {
		elapsed := new(big.Int)
		elapsed = elapsed.Sub(block, &lastHarvested)

		log.Printf("block: %s elapsed: %s required: %d", block.String(), elapsed.String(), uint64(threshold))

		return elapsed.Uint64() >= uint64(threshold), nil
	}
}

func getQuoteForHarvests(
	harvests []strategy.AbstractStrategyHarvest,
	aggregator aggregators.Aggregator,
	chainid uint64,
	stablecoin common.Address,
) (*big.Int, error) {
	amount := new(big.Int)
	errCh := make(chan error)
	amountCh := make(chan *big.Int)

	getQuoteFunc := func(harvest strategy.AbstractStrategyHarvest, errCh chan error, amountCh chan *big.Int) {
		quote, err := aggregator.GetQuote(aggregators.QuoteArgs{
			ChainID: chainid,
			From:    harvest.Token,
			To:      stablecoin,
			Amount:  harvest.Amount,
			Opts:    aggregators.OneInchQuoteOpts{},
		})

		if err != nil {
			errCh <- err
			return
		}

		amountCh <- quote.ToTokenAmount
	}

	for _, harvest := range harvests {
		go getQuoteFunc(harvest, errCh, amountCh)
	}

	for range harvests {
		select {
		case amt := <-amountCh:
			amount = amount.Add(amount, utils.SlippageMillis(amt, 10))
		case <-errCh:
			// we just ignore these as we assume that the aggregate quote won't be enough, and if it is, might be worth selling what we can.
			// fmt.Printf("aggregator.GetQuote err: %s", err)
		}
	}

	return amount, nil
}

func filter(harvests []strategy.AbstractStrategyHarvest, filter []common.Address) (filtered []strategy.AbstractStrategyHarvest) {
	for _, h := range harvests {
		if contains(h.Token, filter) {
			continue
		}

		filtered = append(filtered, h)
	}

	return
}

func calculateThresholdReached(amount *big.Int, tx *types.Transaction, calculator fees.Calculator, decimals int, threshold float64) bool {
	harvestable, _ := new(big.Float).SetInt(amount).Float64()
	harvestable = harvestable / math.Pow(10, float64(decimals))
	if harvestable == 0 {
		return false
	}

	raw, err := calculator.Fee(tx)
	if err != nil {
		log.Printf("failed to calculate fee %s", err.Error())
		return false
	}

	cost, _ := raw.Float64()

	percentage := cost / harvestable

	log.Printf(
		"cost: %f harvestable: %f percentage: %f%% required: %f%%",
		cost,
		harvestable,
		percentage*100,
		threshold*100,
	)

	return percentage <= threshold
}

func contains(address common.Address, filter []common.Address) bool {
	for _, f := range filter {
		if address == f {
			return true
		}
	}

	return false
}

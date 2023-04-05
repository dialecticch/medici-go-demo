package admin

import (
	"fmt"
	"log"
	"math/big"

	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/dialecticch/medici-go/pkg/automator/fees"
	"github.com/dialecticch/medici-go/pkg/automator/sender"
	"github.com/dialecticch/medici-go/pkg/automator/strategy"
	strategy2 "github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/oracles"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/sql"
	"github.com/dialecticch/medici-go/pkg/utils"
	"github.com/ethereum-optimism/optimism/gas-oracle/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	optimism = 10
)

var gasPriceOracleAddress = common.HexToAddress("0x420000000000000000000000000000000000000F")

var (
	claim     bool
	threshold float64
)

var (
	cleanerCmd = &cobra.Command{
		Use:   "cleaner",
		Short: "Checks every strategy for dust harvest and allows claiming them",

		RunE: func(*cobra.Command, []string) error {
			network := Cfg.Networks[chainID]
			conn, err := ethclient.Dial(network.Node)
			if err != nil {
				return errors.Wrap(err, "failed to connect to ethereum client")
			}

			db, err := sql.Open(Cfg.DB)
			if err != nil {
				log.Printf("error opening db: %s", err)
				return err
			}

			strats, err := repositories.NewStrategyRepository(db, chainID).Query(repositories.Inactive)
			if err != nil {
				log.Printf("error fetching strategies: %s", err)
				return err
			}

			stablecoin := common.HexToAddress(network.Stablecoin)
			noSellTokens := utils.StringsToAddrs(network.NoSellTokens)

			aggregator, err := aggregators.NewAggregator(&network, &Cfg.Aggregators, big.NewInt(int64(chainID)))
			if err != nil {
				return errors.Wrap(err, "failed to build aggregator")
			}

			oracle, err := oracles.NewCoinmarketCapPriceOracle(network.Symbol, "USD")
			if err != nil {
				return errors.Wrap(err, "failed to start price oracle")
			}

			var calculator fees.Calculator
			if chainID == optimism {
				gasPrice, err := bindings.NewGasPriceOracle(gasPriceOracleAddress, conn)
				if err != nil {
					return errors.Wrap(err, "failed to create gas price oracle")
				}

				calculator = fees.NewOptimismFeeCalculator(oracle, gasPrice)
			} else {
				calculator = fees.NewEthereumFeeCalculator(oracle)
			}

			privateKey, err := crypto.HexToECDSA(Cfg.PrivateKey)
			if err != nil {
				return errors.Wrap(err, "could not parse private key, is MEDICI_PRIVATE_KEY set?")
			}

			txsender := sender.NewTransactionSender(conn)

			for _, strat := range strats {
				harvests, err := strategy.SimulateClaim(conn, strat.Strategy.Address, strat.Pool, strat.Safe)
				if err != nil {
					return err
				}

				if len(harvests) == 0 {
					continue
				}

				fmt.Printf("Harvests for %s (%s) in pool %s\n", strat.Strategy.Name, strat.Strategy.Address, strat.Pool)
				fmt.Printf("Safe: %s\n", strat.Safe)

				for _, harvest := range harvests {
					fmt.Printf("%s - %s\n", harvest.Token.String(), harvest.Amount.String())
				}

				if !claim {
					fmt.Println()
					continue
				}

				thresholdFunc := strategy.NewGasThresholdFunc(strategy.GasThresholdArguments{
					Backend:         conn,
					StrategyAddress: strat.Strategy.Address,
					ChainID:         chainID,
					Aggregator:      aggregator,
					Calculator:      calculator,
					Decimals:        network.StablecoinDecimals,
					Threshold:       threshold,
					Safe:            strat.Safe,
					StableCoin:      stablecoin,
					NoSellTokens:    noSellTokens,
					Pool:            strat.Pool,
				})

				s, err := strategy2.NewStrategy(strat.Strategy.Address, conn)
				if err != nil {
					return err
				}

				builder := strategy.NewHarvestFunc(strategy.Arguments{
					Backend:         conn,
					Key:             privateKey,
					ChainID:         big.NewInt(int64(chainID)),
					Strategy:        s,
					Safe:            strat.Safe,
					Aggregator:      aggregator,
					StrategyAddress: strat.Strategy.Address,
					StableCoin:      stablecoin,
					NoSellTokens:    noSellTokens,
					Pool:            strat.Pool,
				})

				tx, err := builder()
				if err != nil {
					fmt.Printf("failed to build err %s\n\n", err)
					continue
				}

				reached, err := thresholdFunc(tx)
				if err != nil {
					return err
				}

				if reached {
					_, err := txsender.Send(tx)
					if err != nil {
						log.Printf("txsender.Send err: %s\n\n", err)
					}
				}
			}

			return nil
		},
	}
)

func init() {
	cleanerCmd.Flags().Uint64Var(&chainID, "chainid", 1, "the chainid")
	cleanerCmd.Flags().BoolVar(&claim, "claim", false, "whether to claim harvests or not")
	cleanerCmd.Flags().Float64Var(&threshold, "threshold", 0.5, "the threshold percentage")
}

package control

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/dialecticch/medici-go/pkg/automator/fees"
	"github.com/dialecticch/medici-go/pkg/automator/strategy"
	"github.com/dialecticch/medici-go/pkg/utils"
	"github.com/ethereum-optimism/optimism/gas-oracle/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/dialecticch/medici-go/pkg/automator"
	"github.com/dialecticch/medici-go/pkg/automator/logger"
	"github.com/dialecticch/medici-go/pkg/automator/sender"
	"github.com/dialecticch/medici-go/pkg/oracles"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/sql"
)

const (
	optimism = 10
)

var gasPriceOracleAddress = common.HexToAddress("0x420000000000000000000000000000000000000F")

var (
	automateCmd = &cobra.Command{
		Use:   "automate",
		Short: "Watches new blocks and optionally runs harvest",

		RunE: func(*cobra.Command, []string) error {
			privateKey, err := crypto.HexToECDSA(Cfg.PrivateKey)
			if err != nil {
				return errors.Wrap(err, "could not parse private key, is MEDICI_PRIVATE_KEY set?")
			}

			keyAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
			log.Printf("Loaded private key, address is %s\n", keyAddr)

			db, err := sql.Open(Cfg.DB)
			if err != nil {
				return err
			}

			repo := repositories.NewStrategyRepository(db, ChainID)

			network := Cfg.Networks[ChainID]

			log.Println("Connecting to eth network..")
			conn, err := ethclient.Dial(network.Node)
			if err != nil {
				return errors.Wrap(err, "failed to connect to ethereum client")
			}

			log.Println("Subscribing to new heads..")
			sink := make(chan *types.Header)
			sub, err := conn.SubscribeNewHead(context.Background(), sink)
			if err != nil {
				return errors.Wrap(err, "failed to subscribe to new heads")
			}

			c := make(chan bool)

			go func() {
				err := <-sub.Err()
				if err != nil {
					log.Printf("head subscription err: %s", err)
					c <- true
				}
			}()

			aggregator, err := aggregators.NewAggregator(&network, &Cfg.Aggregators, big.NewInt(int64(ChainID)))
			if err != nil {
				return errors.Wrap(err, "failed to build aggregator")
			}

			oracle, err := oracles.NewCoinmarketCapPriceOracle(network.Symbol, "USD")
			if err != nil {
				return errors.Wrap(err, "failed to start price oracle")
			}

			var calculator fees.Calculator
			if ChainID == optimism {
				gasPrice, err := bindings.NewGasPriceOracle(gasPriceOracleAddress, conn)
				if err != nil {
					return errors.Wrap(err, "failed to create gas price oracle")
				}

				calculator = fees.NewOptimismFeeCalculator(oracle, gasPrice)
			} else {
				calculator = fees.NewEthereumFeeCalculator(oracle)
			}

			l := logger.NewSlackLogger(Cfg.Slack.Webhook, network.Explorer)

			s, err := newTransactionSender(conn, network.SenderRPC)
			if err != nil {
				return err
			}

			builder := strategy.NewBuilder(
				privateKey,
				big.NewInt(int64(ChainID)),
				conn,
				network.StablecoinDecimals,
				common.HexToAddress(network.Stablecoin),
				common.HexToAddress(network.WrappedToken),
				calculator,
				aggregator,
				utils.StringsToAddrs(network.NoSellTokens),
			)

			automator := automator.NewAutomator(
				sink,
				repo,
				builder,
				s,
				privateKey,
				network.Epoch,
				3, // @TODO config
				l,
				conn,
			)

			go func() {
				<-c
				os.Exit(1)
			}()

			return automator.Run()
		},
	}
)

func init() {
	automateCmd.Flags().Uint64Var(&ChainID, "chainid", 1, "the chainid")
}

func newTransactionSender(backend sender.ContractTransactorAndMiner, senderRPC string) (sender.Sender, error) {
	if senderRPC == "" {
		return sender.NewTransactionSender(backend), nil
	}

	conn, err := ethclient.Dial(senderRPC)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to sender RPC")
	}

	return sender.NewTransactionSender(conn), nil
}

package control

import (
	"math/big"

	"github.com/dialecticch/medici-go/pkg/balances"
	"github.com/dialecticch/medici-go/pkg/balances/writers"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/sql"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Watches for balance changes in strategies.",

		RunE: func(*cobra.Command, []string) error {
			db, err := sql.Open(Cfg.DB)
			if err != nil {
				return err
			}

			repo := repositories.NewStrategyRepository(db, ChainID)
			err = repo.Run(repositories.Active)
			if err != nil {
				return err
			}

			network := Cfg.Networks[ChainID]

			conn, err := ethclient.Dial(network.Node)
			if err != nil {
				return errors.Wrap(err, "failed to connect to ethereum client")
			}

			watcher := balances.NewWatcher(repo, conn)
			consumer := balances.NewConsumer(watcher, []writers.Writer{
				writers.NewDatabaseWriter(db, big.NewInt(int64(ChainID))),
				writers.NewSlackWriter(Cfg.Slack.Webhook, network.Explorer),
				writers.NewDepositedAmountWriter(conn, db, big.NewInt(int64(ChainID))),
			})

			return consumer.Run()
		},
	}
)

func init() {
	balancesCmd.Flags().Uint64Var(&ChainID, "chainid", 1, "the chainid")
}

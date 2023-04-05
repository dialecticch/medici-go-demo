package control

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/dialecticch/medici-go/pkg/harvests"
	"github.com/dialecticch/medici-go/pkg/harvests/writers"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/sql"
)

var (
	harvestsCmd = &cobra.Command{
		Use:   "harvests",
		Short: "Watches for and logs any harvest events.",

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

			watcher := harvests.NewWatcher(repo, conn)
			consumer := harvests.NewConsumer(watcher, []writers.Writer{
				writers.NewDatabaseWriter(db, big.NewInt(int64(ChainID))),
				writers.NewLastHarvestedWriter(db, big.NewInt(int64(ChainID))),
				writers.NewSlackWriter(Cfg.Slack.Webhook, network.Explorer),
			})

			return consumer.Run()
		},
	}
)

func init() {
	harvestsCmd.Flags().Uint64Var(&ChainID, "chainid", 1, "the chainid")
}

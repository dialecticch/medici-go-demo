package control

import (
	dbsql "database/sql"

	"github.com/dialecticch/medici-go/pkg/safe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/dialecticch/medici-go/pkg/safe/writers"
	"github.com/dialecticch/medici-go/pkg/sql"
)

var (
	safeCmd = &cobra.Command{
		Use:   "safe",
		Short: "Watches for and logs any safe events.",

		RunE: func(*cobra.Command, []string) error {
			db, err := sql.Open(Cfg.DB)
			if err != nil {
				return err
			}

			network := Cfg.Networks[ChainID]

			conn, err := ethclient.Dial(network.Node)
			if err != nil {
				return errors.Wrap(err, "failed to connect to ethereum client")
			}

			safes, err := fetchSafesForChain(db, ChainID)
			if err != nil {
				return errors.Wrap(err, "failed to fetch safes")
			}

			watcher := safe.NewWatcher(safes, conn)
			consumer := safe.NewConsumer(watcher, []writers.Writer{
				writers.NewSlackWriter(Cfg.Slack.Webhook, network.Explorer),
			})

			return consumer.Run()
		},
	}
)

func init() {
	safeCmd.Flags().Uint64Var(&ChainID, "chainid", 1, "the chainid")
}

func fetchSafesForChain(db *dbsql.DB, chainid uint64) ([]common.Address, error) {
	rows, err := db.Query("SELECT address FROM safes WHERE chain_id = $1", chainid)
	if err != nil {
		return nil, err
	}

	addrs := make([]common.Address, 0)
	for rows.Next() {

		var addr string

		err = rows.Scan(&addr)
		if err != nil {
			return nil, err
		}

		addrs = append(addrs, common.HexToAddress(addr))
	}

	return addrs, nil
}

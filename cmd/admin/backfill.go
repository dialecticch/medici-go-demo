package admin

import (
	"log"
	"math/big"

	"github.com/dialecticch/medici-go/pkg/balances"
	bw "github.com/dialecticch/medici-go/pkg/balances/writers"
	"github.com/dialecticch/medici-go/pkg/harvests"
	hw "github.com/dialecticch/medici-go/pkg/harvests/writers"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/sql"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type UpdateType string

const (
	DEPOSIT  UpdateType = "deposit"
	WITHDRAW UpdateType = "withdraw"
	HARVEST  UpdateType = "harvest"
	ALL      UpdateType = "all"
)

var (
	chainID      uint64
	startBlock   uint64
	backfillType UpdateType
)

var (
	backfillCmd = &cobra.Command{
		Use:   "backfill",
		Short: "Fetches all log events from specified period and ensures events exist in db",

		RunE: func(*cobra.Command, []string) error {
			network := Cfg.Networks[chainID].Node
			conn, err := ethclient.Dial(network)
			if err != nil {
				return errors.Wrap(err, "failed to connect to ethereum client")
			}

			db, err := sql.Open(Cfg.DB)
			if err != nil {
				log.Printf("error opening db: %s", err)
				return err
			}

			strats, err := repositories.NewStrategyRepository(db, chainID).Query(repositories.Any)
			if err != nil {
				log.Printf("error fetching strategies: %s", err)
				return err
			}

			if backfillType == HARVEST || backfillType == ALL {
				log.Println("Backfilling harvests")
				backfiller := harvests.NewBackfiller(conn, []hw.Writer{
					hw.NewDatabaseWriter(db, big.NewInt(int64(chainID))),
					hw.NewLastHarvestedWriter(db, big.NewInt(int64(chainID))),
				})
				err = backfiller.Run(startBlock, strats)
				if err != nil {
					log.Fatalf("backfiller.Run err: %s", err)
				}

				return nil
			}

			log.Println("Backfilling deposits")
			backfiller := balances.NewBackfiller(conn, []bw.Writer{
				bw.NewDatabaseWriter(db, big.NewInt(int64(chainID))),
				bw.NewDepositedAmountWriter(conn, db, big.NewInt(int64(chainID))),
			})

			if backfillType == DEPOSIT || backfillType == ALL {
				err = backfiller.RunDeposits(startBlock, strats)
				if err != nil {
					log.Fatalf("backfiller.RunDeposits err: %s", err)
				}
			}

			if backfillType == WITHDRAW || backfillType == ALL {
				err = backfiller.RunWithdraws(startBlock, strats)
				if err != nil {
					log.Fatalf("backfiller.RunWithdraws err: %s", err)
				}
			}

			return nil
		},
	}
)

func init() {
	backfillCmd.Flags().Uint64Var(&chainID, "chainid", 1, "the chainid")
	backfillCmd.Flags().Uint64Var(&startBlock, "start", 13081142, "block to start backfill from")
	backfillCmd.Flags().StringVar((*string)(&backfillType), "type", "all", "all")
}

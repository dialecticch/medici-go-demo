package pools

import (
	"log"

	"github.com/dialecticch/medici-go/pkg/sql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	poolID uint64
	safeID uint64

	threshold     float64
	action        string
	thresholdType string

	connectCmd = &cobra.Command{
		Use:   "connect",
		Short: "Connects a pool to a safe",

		RunE: func(*cobra.Command, []string) error {
			db, err := sql.Open(Cfg.DB)
			if err != nil {
				return err
			}

			if thresholdType == "blocks" {
				thresholdType = "elapsed_blocks"
			} else {
				thresholdType = "gas_percentage"
			}

			stmt, err := db.Prepare("INSERT INTO safe_pools (safe_id, pool_id, threshold, action, threshold_type) VALUES ($1, $2, $3, $4, $5)")
			if err != nil {
				return errors.Wrap(err, "failed to prepare")
			}

			_, err = stmt.Exec(safeID, poolID, threshold, action, thresholdType)
			if err != nil {
				return errors.Wrap(err, "failed to insert")
			}

			log.Printf(
				"Added Pool %d to Safe %d with threshold %s (%f) and action %s",
				poolID,
				safeID,
				thresholdType,
				threshold,
				action,
			)

			return nil
		},
	}
)

func init() {
	connectCmd.Flags().Uint64Var(&poolID, "pool", 0, "the pool id")
	connectCmd.Flags().Uint64Var(&safeID, "safe", 0, "the safe id")
	connectCmd.Flags().Float64Var(&threshold, "threshold", 0.05, "the threshold")
	connectCmd.Flags().StringVar(&action, "action", "harvest", "the action (harvest or compound)")
	connectCmd.Flags().StringVar(&thresholdType, "type", "gas", "the threshold type (gas or blocks)")
}

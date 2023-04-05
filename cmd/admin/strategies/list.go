package strategies

import (
	"fmt"

	"github.com/dialecticch/medici-go/pkg/sql"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all strategies in the medici database",

		RunE: func(*cobra.Command, []string) error {
			db, err := sql.Open(Cfg.DB)
			if err != nil {
				return err
			}

			query := "SELECT id, chain_id, name FROM strategies"

			args := make([]any, 0)
			if chainID != 0 {
				query += " WHERE chain_id = ?"
				args = append(args, chainID)
			}

			query += " ORDER BY id ASC"

			res, err := db.Query(query, args...)
			if err != nil {
				return err
			}

			for res.Next() {
				var id uint64
				var chain uint64
				var name string

				err = res.Scan(&id, &chain, &name)
				if err != nil {
					return err
				}

				fmt.Printf("%d) %s (%d)\n", id, name, chain)

				rows, err := db.Query("SELECT address, version FROM versions WHERE strategy_id = $1 ORDER BY version DESC", id)
				if err != nil {
					fmt.Println("failed to fetch versions")
				}

				fmt.Println("  Versions:")
				for rows.Next() {
					var addr string
					var version string

					err = rows.Scan(&addr, &version)
					if err != nil {
						continue
					}

					fmt.Printf("    - %s (%s)\n", addr, version)
				}

				rows, err = db.Query("SELECT pool, name FROM pools WHERE strategy_id = $1", id)
				if err != nil {
					fmt.Println("failed to fetch pools")
					continue
				}

				fmt.Println("  Pools:")
				for rows.Next() {
					var pool string
					var name string

					err = rows.Scan(&pool, &name)
					if err != nil {
						continue
					}

					fmt.Printf("    - %s (%s)\n", name, pool)
				}

				fmt.Println()
			}

			return nil
		},
	}
)

func init() {
	listCmd.Flags().Uint64Var(&chainID, "chainid", 0, "the chainid")
}

package pools

import (
	dbsql "database/sql"
	"log"
	"math/big"
	"strings"

	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/importers"
	"github.com/dialecticch/medici-go/pkg/sql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	rawPool    string
	strategyID uint64

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds a pool to the medici database",

		RunE: func(*cobra.Command, []string) error {
			if rawPool == "" {
				return errors.New("valid address required")
			}

			pool, ok := new(big.Int).SetString(rawPool, 10)
			if !ok {
				return errors.New("invalid pool")
			}

			db, err := sql.Open(Cfg.DB)
			if err != nil {
				return err
			}

			network := Cfg.Networks[chainID]

			log.Println("Connecting to eth network..")
			conn, err := ethclient.Dial(network.Node)
			if err != nil {
				return errors.Wrap(err, "failed to connect to ethereum client")
			}

			row := db.QueryRow("SELECT address FROM latest_versions WHERE strategy_id = $1", strategyID)

			var addr string
			err = row.Scan(&addr)

			if err != nil {
				return errors.Wrap(err, "failed to load strategy")
			}

			s, err := strategy.NewStrategyCaller(common.HexToAddress(addr), conn)
			if err != nil {
				return errors.Wrap(err, "failed to connect to contract")
			}

			name, err := s.PoolName(nil, pool)
			if err != nil {
				return errors.Wrap(err, "failed read s.NAME()")
			}

			token, err := s.DepositToken(nil, pool)
			if err != nil {
				return errors.Wrap(err, "failed read s.DepositToken()")
			}

			depositToken, err := getOrInsertToken(token, chainID, db, conn)
			if err != nil {
				return err
			}

			row = db.QueryRow(
				"INSERT INTO pools (pool, name, strategy_id, token_id) VALUES ($1, $2, $3, $4) RETURNING id",
				pool.Int64(),
				name,
				strategyID,
				depositToken,
			)

			var id int
			err = row.Scan(&id)
			if err != nil {
				return errors.Wrap(err, "failed to get id")
			}

			log.Printf("Added pool %s with id %d", name, id)

			return nil
		},
	}
)

func init() {
	addCmd.Flags().StringVar(&rawPool, "pool", "", "the pool id")
	addCmd.Flags().Uint64Var(&strategyID, "strategy", 0, "the id of the strategy")
	addCmd.Flags().Uint64Var(&chainID, "chainid", 1, "the chainid")
}

func getOrInsertToken(address common.Address, chainid uint64, db *dbsql.DB, client *ethclient.Client) (int, error) {
	row := db.QueryRow("SELECT id FROM tokens WHERE address = $1 AND chain_id = $2", strings.ToLower(address.String()), chainid)

	var id int
	err := row.Scan(&id)
	if err == nil {
		return id, nil
	}

	if err != nil && err != dbsql.ErrNoRows {
		return 0, err
	}

	return importers.NewTokenImporter(db, client).Import(address)
}

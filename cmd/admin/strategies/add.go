package strategies

import (
	dbsql "database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/sql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	address string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds a strategy to the medici database",

		RunE: func(*cobra.Command, []string) error {
			if address == "" {
				return errors.New("valid address required")
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

			addr := common.HexToAddress(address)

			id, err := insertOrFetchStrategy(db, chainID, conn, addr)
			if err != nil {
				return err
			}

			return insertVersion(db, conn, id, addr)
		},
	}
)

func insertOrFetchStrategy(db *dbsql.DB, chainid uint64, client *ethclient.Client, addr common.Address) (int, error) {
	m, err := strategy.NewStrategyCaller(common.HexToAddress(address), client)
	if err != nil {
		return 0, errors.Wrap(err, "failed to connect to contract")
	}

	name, err := m.NAME(nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed read m.NAME()")
	}

	row := db.QueryRow("SELECT id FROM strategies WHERE chain_id = $1 AND name = $2", chainid, name)

	var id int
	err = row.Scan(&id)
	if err == nil {
		log.Printf("Fetched strategy %s with id %d", name, id)
		return id, nil
	}

	if err != dbsql.ErrNoRows {
		return 0, errors.Wrap(err, "row.Scan failed")
	}

	row = db.QueryRow(
		"INSERT INTO strategies (chain_id, name) VALUES ($1, $2) RETURNING id",
		chainid,
		name,
	)

	err = row.Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get id")
	}

	fmt.Printf("Inserted strategy %s with id %d", name, id)

	return id, nil
}

func insertVersion(db *dbsql.DB, client *ethclient.Client, strategyID int, addr common.Address) error {
	m, err := strategy.NewStrategyCaller(common.HexToAddress(address), client)
	if err != nil {
		return errors.Wrap(err, "failed to connect to contract")
	}

	version, err := m.VERSION(nil)
	if err != nil {
		return errors.Wrap(err, "failed read m.VERSION()")
	}

	stmt, err := db.Prepare("INSERT INTO versions (strategy_id, address, version) VALUES ($1, $2, $3)")
	if err != nil {
		return errors.Wrap(err, "db.Prepare failed")
	}

	_, err = stmt.Exec(strategyID, strings.ToLower(addr.String()), version)
	if err != nil {
		return err
	}

	log.Printf("Inserted new version %s", version)

	return nil
}

func init() {
	addCmd.Flags().StringVar(&address, "address", "", "the address of the strategy")
	addCmd.Flags().Uint64Var(&chainID, "chainid", 1, "the chainid")
}

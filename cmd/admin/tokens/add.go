package tokens

import (
	"log"

	"github.com/dialecticch/medici-go/pkg/importers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/dialecticch/medici-go/pkg/sql"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var (
	address string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds a token to the medici database",

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

			importer := importers.NewTokenImporter(db, conn)

			id, err := importer.Import(common.HexToAddress(address))
			if err != nil {
				return err
			}

			log.Printf("Imported new token with id %d", id)

			return nil
		},
	}
)

func init() {
	addCmd.Flags().StringVar(&address, "address", "", "the address of the token")
	addCmd.Flags().Uint64Var(&chainID, "chainid", 1, "the chainid")
}

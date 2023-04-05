package importers

import (
	"context"
	"database/sql"
	"strings"

	"github.com/dialecticch/medici-go/pkg/contracts/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type TokenImporter struct {
	db     *sql.DB
	client *ethclient.Client
}

func NewTokenImporter(db *sql.DB, client *ethclient.Client) *TokenImporter {
	return &TokenImporter{
		db:     db,
		client: client,
	}
}

func (t *TokenImporter) Import(addr common.Address) (int, error) {
	caller, err := token.NewToken(addr, t.client)
	if err != nil {
		return 0, errors.Wrap(err, "failed to connect to contract")
	}

	name, err := caller.Name(nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed read t.Name()")
	}

	symbol, err := caller.Symbol(nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed read t.Symbol()")
	}

	decimals, err := caller.Decimals(nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed read t.StablecoinDecimals()")
	}

	chainid, err := t.client.ChainID(context.TODO())
	if err != nil {
		return 0, errors.Wrap(err, "failed to read t.client.ChainID()")
	}

	row := t.db.QueryRow(
		"INSERT INTO tokens (address, name, symbol, decimals, chain_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		strings.ToLower(addr.String()),
		name,
		symbol,
		decimals,
		chainid.Int64(),
	)

	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get id")
	}

	return id, nil
}

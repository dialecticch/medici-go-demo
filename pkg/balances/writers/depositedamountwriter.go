package writers

import (
	"database/sql"
	"math/big"
	"strings"

	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type DepositedAmountWriter struct {
	backend bind.ContractBackend
	db      *sql.DB
	chainid *big.Int
}

func NewDepositedAmountWriter(backend bind.ContractBackend, db *sql.DB, chainid *big.Int) *DepositedAmountWriter {
	return &DepositedAmountWriter{backend: backend, db: db, chainid: chainid}
}

func (d *DepositedAmountWriter) Write(event *Event) error {
	m, err := strategy.NewStrategyCaller(event.Strategy, d.backend)
	if err != nil {
		return err
	}

	deposited, err := m.DepositedAmount(nil, event.Pool, event.Safe)
	if err != nil {
		return err
	}

	stmt, err := d.db.Prepare(`
		UPDATE safe_pools 
		SET deposited_amount = $1, last_deposited_amount_update = $2
		WHERE safe_id = (
			SELECT id FROM safes
			WHERE address = $3 AND chain_id = $6
		) AND pool_id = (
			SELECT id FROM pools
			WHERE pool = $4 AND strategy_id = (SELECT id FROM strategies WHERE address = $5 AND chain_id = $6)
		) AND last_deposited_amount_update < $2`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		deposited.String(),
		event.Block,
		strings.ToLower(event.Safe.String()),
		event.Pool.String(),
		strings.ToLower(event.Strategy.String()),
		d.chainid.Int64(),
	)

	return err
}

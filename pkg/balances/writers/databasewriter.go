package writers

import (
	"database/sql"
	"math/big"
	"strings"
)

type DatabaseWriter struct {
	db      *sql.DB
	chainid *big.Int
}

func NewDatabaseWriter(db *sql.DB, chainid *big.Int) *DatabaseWriter {
	return &DatabaseWriter{
		db:      db,
		chainid: chainid,
	}
}

func (d *DatabaseWriter) Write(event *Event) error {
	stmt, err := d.db.Prepare(`
		INSERT INTO transactions (timestamp, transaction_type, tx, log_index, amount, pool_safe_id) 
		VALUES (
		        $1,
		        $2,
		        $3,
		        $4,
		        $5, 
		        (
		            SELECT id FROM safe_pools 
		            	WHERE safe_id = (SELECT id FROM safes WHERE address = $6 AND chain_id = $9)
		            	AND pool_id = (
		            	    SELECT id FROM pools WHERE pool = $7 
						   	AND strategy_id = (SELECT strategy_id from versions WHERE address = $7)
						)
				)
			)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		event.UpdatedAt,
		event.UpdateType,
		strings.ToLower(event.Transaction.String()),
		event.LogIndex,
		event.Amount.String(),
		strings.ToLower(event.Safe.String()),
		event.Pool.String(),
		strings.ToLower(event.Strategy.String()),
		d.chainid.Int64(),
	)
	return err
}

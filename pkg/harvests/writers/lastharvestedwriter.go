package writers

import (
	"database/sql"
	"math/big"
	"strings"
)

type LastHarvestedWriter struct {
	db      *sql.DB
	chainid *big.Int
}

func NewLastHarvestedWriter(db *sql.DB, chainid *big.Int) *LastHarvestedWriter {
	return &LastHarvestedWriter{
		db:      db,
		chainid: chainid,
	}
}

func (l *LastHarvestedWriter) Write(event *Event) error {
	stmt, err := l.db.Prepare(`
		UPDATE safe_pools 
		SET last_harvested = $1
		WHERE safe_id = (
			SELECT id FROM safes
			WHERE address = $2 AND chain_id = $5
		) AND pool_id = (
			SELECT id FROM pools
			WHERE pool = $3 AND strategy_id = (SELECT strategy_id from versions WHERE address = $4)
		) AND last_harvested < $1`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		event.Block,
		strings.ToLower(event.Safe.String()),
		event.Pool.String(),
		strings.ToLower(event.Strategy.String()),
		l.chainid.Int64(),
	)
	return err
}

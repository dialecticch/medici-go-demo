// Package sql provides helper functions for sql database code.
package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/dialecticch/medici-go/pkg/config"
)

// Open opens a Postgres database from the passed config.
func Open(config config.PostgresConf) (*sql.DB, error) {
	return sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.Database, config.SSL,
		),
	)
}

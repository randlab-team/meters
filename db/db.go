package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	driver = "postgres"

	maxOpenConnections    = 25
	maxIdleConnections    = 25
	maxConnectionLifetime = 5 * time.Minute
)

func InitDB(dbString string) (*sqlx.DB, error) {
	dbConn, err := sqlx.Connect(driver, dbString)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to the db")

		return nil, errors.Wrap(err, "failed to connect to the db")
	}

	dbConn.SetMaxOpenConns(maxOpenConnections)
	dbConn.SetMaxIdleConns(maxIdleConnections)
	dbConn.SetConnMaxLifetime(maxConnectionLifetime)

	if err = dbConn.Ping(); err != nil {
		log.Error().Err(err).Msg("failed to ping database")

		return nil, errors.Wrap(err, "failed to ping database")
	}

	return dbConn, nil
}

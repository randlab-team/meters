package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/randlab-team/meters/models"
)

const (
	createQuery = `
		INSERT INTO meters (
			sn, correct, param_name, "index", date_register, "value", log_interval, status
		) 
        VALUES (
            :sn, :correct, :param_name, :index,:date_register,:value,:log_interval,:status
        )`

	getAllQuery = `SELECT * FROM meters;`
)

//go:generate mockgen -destination=../mocks/repository.go . Meters
type Meters interface {
	Save(meterLog models.MeterLog) error
	GetAll() ([]models.MeterLog, error)
}

type MetersRepo struct {
	db *sqlx.DB

	createStmt *sqlx.NamedStmt
	getAllStmt *sqlx.Stmt
}

func NewMeters(db *sqlx.DB) (*MetersRepo, error) {
	createStmt, err := db.PrepareNamed(createQuery)
	if err != nil {
		log.Error().Err(err).Msg("failed to prepare crate meter log statement")

		return nil, errors.Wrap(err, "failed to prepare crate meter log statement")
	}

	getAllStmt, err := db.Preparex(getAllQuery)
	if err != nil {
		log.Error().Err(err).Msgf("failed to prepare get all meter logs statement")

		return nil, errors.Wrap(err, "failed to prepare get all meter logs statement")
	}

	return &MetersRepo{
		db:         db,
		createStmt: createStmt,
		getAllStmt: getAllStmt,
	}, nil
}

func (m MetersRepo) Save(meterLog models.MeterLog) error {
	if _, err := m.createStmt.Exec(meterLog); err != nil {
		log.Fatal().Err(err).Msg("failed to save meter log")

		return errors.Wrap(err, "failed to save meter log")
	}

	return nil
}

func (m MetersRepo) GetAll() ([]models.MeterLog, error) {
	var meterLogs []models.MeterLog
	if err := m.getAllStmt.Select(&meterLogs); err != nil {
		log.Fatal().Err(err).Msg("failed to fetch all meter logs")

		return []models.MeterLog{}, errors.Wrap(err, "failed to fetch all meter logs")
	}

	return meterLogs, nil
}

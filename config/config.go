package config

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	dbStringKey = "db_string"
)

type AppConfig struct {
	DbString string `mapstructure:"db_string"`
}

func InitConfig() (AppConfig, error) {
	if err := initEnv(); err != nil {
		log.Error().Err(err).Msg("failed to initialize config env")

		return AppConfig{}, errors.Wrap(err, "failed to initialize config env")
	}

	appConf := AppConfig{}
	err := viper.Unmarshal(&appConf)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch config form env")

		return AppConfig{}, errors.Wrap(err, "failed to fetch config form env")
	}

	return appConf, nil
}

func initEnv() error {
	err := viper.BindEnv(dbStringKey)
	if err != nil {
		log.Error().Err(err).Msgf("failed to bind env var: %s", dbStringKey)

		return errors.Wrapf(err, "failed to bind env var: %s", dbStringKey)
	}

	return nil
}

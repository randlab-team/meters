package main

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/randlab-team/meters/db"
	"github.com/randlab-team/meters/models"
	"github.com/randlab-team/meters/repository"
)

const (
	MeterLogsCsvFilePath = "csv_path"

	csvSeparator = ';'
)

func init() {
	rootCmd.PersistentFlags().StringP(MeterLogsCsvFilePath, "f", "", "Provide csv file with meter logs")
	viper.BindPFlag(MeterLogsCsvFilePath, rootCmd.PersistentFlags().Lookup(MeterLogsCsvFilePath))

	rootCmd.AddCommand(csvCmd)
}

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "Csv importer for meter logs",
	Run: func(cmd *cobra.Command, args []string) {
		initLogger()
		initCsvReader()

		dbString := viper.GetString(DbStringKey)
		csvFilePath := viper.GetString(MeterLogsCsvFilePath)

		dbConn := initDb(dbString)
		meterLogRepo, err := repository.NewMeters(dbConn)
		if err != nil {
			log.Fatal().Err(err).Msg("filed to init meters repo")
		}

		metersCsvFile, err := os.OpenFile(csvFilePath, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to open meter logs csv file")
		}
		defer metersCsvFile.Close()

		var meterCsvLogs []models.MeterCsvLog
		if err := gocsv.UnmarshalFile(metersCsvFile, &meterCsvLogs); err != nil {
			log.Fatal().Err(err).Msg("filed to unmarshal meter logs csv")
		}

		for _, meterCsvLog := range meterCsvLogs {
			meterLog, err := meterCsvLog.ToMeterLog()
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to map meter log from csv")
			}

			err = meterLogRepo.Save(meterLog)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to save meter log")
			}

			log.Info().Msgf("properly saved meter log: %+v", meterCsvLog)
		}
		log.Info().Msgf("properly processed meter logs")
	},
}

func initLogger() {
	log.Logger = log.With().Caller().Logger()
}

func initCsvReader() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		csvReader := csv.NewReader(in)
		csvReader.Comma = csvSeparator

		return csvReader
	})
}

func initDb(dbString string) *sqlx.DB {
	dbConn, err := db.InitDB(dbString)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	return dbConn
}

package models

import (
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	CsvCorrectTrueValue = "Tak"
	CsvDateLayout       = "01.02.2006 15:04"
)

type MeterCsvLog struct {
	SN           int32  `csv:"SN"`
	Correct      string `csv:"Poprawna"`
	ParamName    string `csv:"Nazwa parametru"`
	Index        int64  `csv:"Indeks"`
	DateRegister string `csv:"Data rejestracji w urzędzeniu"`
	Value        int64  `csv:"Wartość [m3]"`
	LogInterval  int32  `csv:"Interwał rejestracji"`
	Status       int32  `csv:"Status"`
}

func (ml MeterCsvLog) GetCorrect() bool {
	return ml.Correct == CsvCorrectTrueValue
}

func (ml MeterCsvLog) GetDateRegister() (time.Time, error) {
	dateReg, err := time.Parse(CsvDateLayout, ml.DateRegister)
	if err != nil {
		log.Error().Err(err).Msgf("failed to parse date form csv: %s with format: %s", ml.DateRegister, CsvDateLayout)

		return time.Time{}, errors.Wrapf(err, "failed to parse date form csv: %s with format: %s", ml.DateRegister, CsvDateLayout)
	}

	return dateReg, nil
}

func (ml MeterCsvLog) ToMeterLog() (MeterLog, error) {
	dateRegister, err := ml.GetDateRegister()
	if err != nil {
		return MeterLog{}, errors.Wrap(err, "failed to parse date from meter log csv")
	}

	return MeterLog{
		SN:           ml.SN,
		Correct:      ml.GetCorrect(),
		ParamName:    ml.ParamName,
		Index:        ml.Index,
		DateRegister: dateRegister,
		Value:        ml.Value,
		LogInterval:  ml.LogInterval,
		Status:       ml.Status,
	}, nil
}

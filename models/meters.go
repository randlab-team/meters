package models

import "time"

type MeterLog struct {
	ID           int64     `json:"id"            db:"id"`
	SN           int32     `json:"sn"            db:"sn"`
	Correct      bool      `json:"correct"       db:"correct"`
	ParamName    string    `json:"param_name"    db:"param_name"`
	Index        int64     `json:"index"         db:"index"`
	DateRegister time.Time `json:"date_register" db:"date_register"`
	Value        int64     `json:"value"         db:"value"`
	LogInterval  int32     `json:"log_interval"  db:"log_interval"`
	Status       int32     `json:"status"        db:"status"`
}

type GetMererLogRequest struct {
	ID uint `query:"id"`
}

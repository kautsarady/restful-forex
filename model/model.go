package model

import (
	"time"
)

// Exchange .
type Exchange struct {
	From  string    `json:"from" example:"USD"`
	To    string    `json:"to" example:"IDR"`
	Rate  float64   `json:"rate" example:"14967.667"`
	Rate7 float64   `json:"rate7" example:"0.233"`
	Date  time.Time `json:"date" example:"2016-02-01"`
}

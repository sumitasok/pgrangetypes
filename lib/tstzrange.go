package lib

import (
	"database/sql/driver"
	"time"
)

type TstzrangeI interface {
	fromTimeString() string
	toTimeString() string
	Scan(src interface{}) error
	Value() (driver.Value, error)
}

var _ TstzrangeI = Tstzrange{}

func NewTstzrange(prefix rune, fromTime, toTime time.Time, postfix rune) (*Tstzrange, error) {
	return &Tstzrange{prefix: prefix, fromTime: fromTime, toTime: toTime, postfix: postfix}, nil
}

type Tstzrange struct {
	prefix   rune
	fromTime time.Time
	toTime   time.Time
	postfix  rune
}

func (t Tstzrange) fromTimeString() string {
	return t.fromTime.Format("2006-01-02 15:04:05-07:00")
}

func (t Tstzrange) toTimeString() string {
	return t.toTime.Format("2006-01-02 15:04:05-07:00")
}

func (t Tstzrange) Scan(src interface{}) error {
	return nil
}

func (t Tstzrange) Value() (driver.Value, error) {
	//TODO: check if from is greater than to; return error.
	return "[" + t.fromTimeString() + "," + t.toTimeString() + ")", nil
}

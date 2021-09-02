package lib

import (
	"database/sql/driver"
	"time"
)

type tstzrange struct {
	prefix   rune
	fromTime time.Time
	toTime   time.Time
	postfix  rune
}

func (t tstzrange) fromTimeString() string {
	return t.fromTime.Format("2006-01-02 15:04:05-07:00")
}

func (t tstzrange) toTimeString() string {
	return t.toTime.Format("2006-01-02 15:04:05-07:00")
}

func (t tstzrange) Scan(src interface{}) error {
	return nil
}

func (t tstzrange) Value() (driver.Value, error) {
	//TODO: check if from is greater than to; return error.
	return "[" + t.fromTimeString() + "," + t.toTimeString() + ")", nil
}

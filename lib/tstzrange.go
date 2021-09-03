package lib

import (
	"database/sql/driver"
	"strings"
	"time"
)

type TstzrangeI interface {
	fromTimeString() string
	toTimeString() string
	Scan(src interface{}) error
	Value() (driver.Value, error)
	ToString() string
}

//var _ TstzrangeI = Tstzrange{}
var timeFormat string = "2006-01-02 15:04:05-07:00"

func NewTstzrange(prefix rune, fromTime, toTime time.Time, postfix rune) (*Tstzrange, error) {
	return &Tstzrange{prefix: prefix, fromTime: fromTime, toTime: toTime, postfix: postfix}, nil
}

type Tstzrange struct {
	prefix   rune
	fromTime time.Time
	toTime   time.Time
	postfix  rune
}

func (t Tstzrange) ToString() string {
	return string(t.prefix) + t.fromTimeString() + "," + t.toTimeString() + string(t.postfix)
}

func (t Tstzrange) fromTimeString() string {
	return t.fromTime.Format(timeFormat)
}

func (t Tstzrange) toTimeString() string {
	return t.toTime.Format(timeFormat)
}

func (t *Tstzrange) Scan(src interface{}) error {
	str := src.(string)
	t.prefix = rune(str[0])
	t.postfix = rune(str[len(str)-1])
	str = strings.Trim(str, "[]()\"")

	fromTo := strings.Split(str, ",")

	from := strings.Trim(fromTo[0], "\"")
	fromTime, err := time.Parse(timeFormat, from)
	if err != nil {
		return err
	}

	to := strings.Trim(fromTo[1], "\"")
	toTime, err := time.Parse(timeFormat, to)
	if err != nil {
		return err
	}

	t.fromTime = fromTime
	t.toTime = toTime

	return nil
}

func (t Tstzrange) Value() (driver.Value, error) {
	//TODO: check if from is greater than to; return error.
	return "[" + t.fromTimeString() + "," + t.toTimeString() + ")", nil
}

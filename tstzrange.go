package awesomeProject

import (
	"database/sql/driver"
	"errors"
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
var timeFormat = "2006-01-02 15:04:05-07:00"

func NewTstzrange(prefix rune, fromTime, toTime time.Time, postfix rune) (*Tstzrange, error) {
	return &Tstzrange{prefix: prefix, FromTime: DateParser{fromTime}, ToTime: DateParser{toTime}, postfix: postfix}, nil
}

type Tstzrange struct {
	prefix   rune
	FromTime DateParser `json:"from_time"`
	ToTime   DateParser `json:"to_time"`
	postfix  rune
}

func (t Tstzrange) ToString() string {
	prefix := string(t.prefix)
	if t.prefix == 0 {
		prefix = "["
	} // default prefix

	postfix := string(t.postfix)
	if t.postfix == 0 {
		postfix = ")"
	} // default postfix

	return prefix + t.fromTimeString() + "," + t.toTimeString() + postfix
}

func (t Tstzrange) fromTimeString() string {
	return t.FromTime.Format(timeFormat)
}

func (t Tstzrange) toTimeString() string {
	return t.ToTime.Format(timeFormat)
}

func (t *Tstzrange) Scan(src interface{}) error {
	str := src.(string)
	//TODO: validations
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

	t.FromTime = DateParser{fromTime}
	t.ToTime = DateParser{toTime}

	return nil
}

func (t Tstzrange) Value() (driver.Value, error) {
	if t.FromTime.After(t.ToTime.Time) {
		return nil, errors.New("from time cannot be after to time")
	}

	return t.ToString(), nil
}
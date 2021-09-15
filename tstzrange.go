package pgrangetypes

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
var timeFormat = "2006-01-02T15:04:05-07:00"

func NewTstzrange(prefix rune, fromTime, toTime time.Time, postfix rune) (*Tstzrange, error) {
	return &Tstzrange{prefix: prefix, FromTime: DateParser{fromTime}, ToTime: DateParser{toTime}, postfix: postfix}, nil
}

type Tstzrange struct {
	prefix   rune
	FromTime DateParser `json:"fromTime"`
	ToTime   DateParser `json:"toTime"`
	postfix  rune
}

func (t Tstzrange) String() string {
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

func (t Tstzrange) Empty() bool {
	if t.FromTime.Equal(time.Time{}) {
		return true
	}

	if t.ToTime.Equal(time.Time{}) {
		return true
	}

	return false
}

func (t *Tstzrange) Scan(src interface{}) error {
	str := src.(string)

	if str == "empty" {
		t.prefix = '['
		t.FromTime = DateParser{time.Time{}}
		t.ToTime = DateParser{time.Time{}}
		t.postfix = ')'
		return nil
	}
	//TODO: validations
	t.prefix = rune(str[0])
	t.postfix = rune(str[len(str)-1])
	str = strings.Trim(str, "[]()\"")

	fromTo := strings.Split(str, ",")

	from := strings.Trim(fromTo[0], "\"") + ":00"
	from = strings.Replace(from, " ", "T", 1)
	fromTime, err := time.Parse(timeFormat, from)
	if err != nil {
		return err
	}

	to := strings.Trim(fromTo[1], "\"") + ":00"
	to = strings.Replace(to, " ", "T", 1)
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

	return t.String(), nil
}

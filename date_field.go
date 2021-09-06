package pgrangetypes

import (
	"fmt"
	"time"
)

// https://stackoverflow.com/questions/54618633/parsing-time-as-2006-01-02t150405z0700-cannot-parse-as-2006
// date format // Mon, 02 Jan 2006 15:04:05 -0700
type DateParser struct {
	time.Time
}

func (df *DateParser) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"`+time.RFC1123Z+`"`, string(data))
	*df = DateParser{tt}
	return err
}

func (df DateParser) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(df.Time).Format(time.RFC1123Z))
	return []byte(stamp), nil
}

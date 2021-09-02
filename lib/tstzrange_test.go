package lib

import (
	"database/sql/driver"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func Test_tstzrange_Scan(t1 *testing.T) {
	type fields struct {
		prefix   rune
		fromTime time.Time
		toTime   time.Time
		postfix  rune
	}
	type args struct {
		src interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tstzrange{
				prefix:   tt.fields.prefix,
				fromTime: tt.fields.fromTime,
				toTime:   tt.fields.toTime,
				postfix:  tt.fields.postfix,
			}
			if err := t.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t1.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tstzrange_Value(t1 *testing.T) {
	layout := "2006-01-02T15:04:05-07:00"
	str := "2014-11-12T11:45:26+05:30"
	timeExample, err := time.Parse(layout, str)
	assert := assert.New(t1)
	assert.NoError(err)

	type fields struct {
		prefix   rune
		fromTime time.Time
		toTime   time.Time
		postfix  rune
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "ValidValue",
			fields: fields{
				prefix:   '[',
				fromTime: timeExample,
				toTime:   timeExample.Add(time.Duration(1 * time.Hour)),
				postfix:  ')',
			},
			want:    "[2014-11-12 11:45:26+05:30,2014-11-12 12:45:26+05:30)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tstzrange{
				prefix:   tt.fields.prefix,
				fromTime: tt.fields.fromTime,
				toTime:   tt.fields.toTime,
				postfix:  tt.fields.postfix,
			}
			got, err := t.Value()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tstzrange_fromTimeString(t1 *testing.T) {
	layout := "2006-01-02T15:04:05-07:00"
	str := "2014-11-12T11:45:26+05:30"
	timeExample, err := time.Parse(layout, str)
	assert := assert.New(t1)
	assert.NoError(err)

	type fields struct {
		prefix   rune
		fromTime time.Time
		toTime   time.Time
		postfix  rune
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ConvertToString",
			fields: fields{
				prefix:   '[',
				fromTime: timeExample,
				toTime:   timeExample.Add(time.Duration(1 * time.Hour)),
				postfix:  ')',
			},
			want: "2014-11-12 11:45:26+05:30",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tstzrange{
				prefix:   tt.fields.prefix,
				fromTime: tt.fields.fromTime,
				toTime:   tt.fields.toTime,
				postfix:  tt.fields.postfix,
			}
			if got := t.fromTimeString(); got != tt.want {
				t1.Errorf("fromTimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tstzrange_toTimeString(t1 *testing.T) {
	layout := "2006-01-02T15:04:05-07:00"
	str := "2014-11-12T11:45:26+05:30"
	timeExample, err := time.Parse(layout, str)
	assert := assert.New(t1)
	assert.NoError(err)

	type fields struct {
		prefix   rune
		fromTime time.Time
		toTime   time.Time
		postfix  rune
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ConvertToString",
			fields: fields{
				prefix:   '[',
				fromTime: timeExample,
				toTime:   timeExample.Add(time.Duration(1 * time.Hour)),
				postfix:  ')',
			},
			want: "2014-11-12 12:45:26+05:30",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tstzrange{
				prefix:   tt.fields.prefix,
				fromTime: tt.fields.fromTime,
				toTime:   tt.fields.toTime,
				postfix:  tt.fields.postfix,
			}
			if got := t.toTimeString(); got != tt.want {
				t1.Errorf("toTimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

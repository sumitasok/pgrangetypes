package lib

import (
	"database/sql/driver"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func Test_Tstzrange_Scan(t1 *testing.T) {
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
	type args struct {
		src interface{}
	}

	_fields := fields{
		prefix:   '[',
		fromTime: timeExample,
		toTime:   timeExample.Add(time.Duration(1 * time.Hour)),
		postfix:  ')',
	}
	_tstzrange, err := NewTstzrange(_fields.prefix, _fields.fromTime, _fields.toTime, _fields.postfix)
	assert.NoError(err)

	type want struct {
		prefix  rune
		postfix rune
		fields  fields
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    want
	}{
		{
			name:    "ValidScan",
			fields:  _fields,
			args:    args{src: string(_fields.prefix) + _tstzrange.fromTimeString() + "," + _tstzrange.toTimeString() + string(_fields.postfix)},
			wantErr: false,
			want: want{
				prefix:  _fields.prefix,
				postfix: _fields.postfix,
				fields:  _fields,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tstzrange{
				prefix:   tt.fields.prefix,
				FromTime: tt.fields.fromTime,
				ToTime:   tt.fields.toTime,
				postfix:  tt.fields.postfix,
			}
			if err := t.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t1.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}

			if t.prefix != tt.want.prefix {
				t1.Errorf("Scan() want = %v, got %v", string(tt.want.prefix), string(t.prefix))
			}

			if t.postfix != tt.want.postfix {
				t1.Errorf("Scan() want = %v, got %v", string(tt.want.postfix), string(t.postfix))
			}

			if t.FromTime != tt.want.fields.fromTime {
				t1.Errorf("Scan() want = %v, got %v", tt.want.fields.fromTime, t.FromTime)
			}
		})
	}
}

func Test_Tstzrange_Value(t1 *testing.T) {
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
		{
			name: "InValidTime_From_after_To",
			fields: fields{
				prefix:   '[',
				fromTime: timeExample,
				toTime:   timeExample.Add(time.Duration(-1 * time.Hour)),
				postfix:  ')',
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tstzrange{
				prefix:   tt.fields.prefix,
				FromTime: tt.fields.fromTime,
				ToTime:   tt.fields.toTime,
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

func Test_Tstzrange_fromTimeString(t1 *testing.T) {
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
			t := Tstzrange{
				prefix:   tt.fields.prefix,
				FromTime: tt.fields.fromTime,
				ToTime:   tt.fields.toTime,
				postfix:  tt.fields.postfix,
			}
			if got := t.fromTimeString(); got != tt.want {
				t1.Errorf("fromTimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Tstzrange_toTimeString(t1 *testing.T) {
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
			t := Tstzrange{
				prefix:   tt.fields.prefix,
				FromTime: tt.fields.fromTime,
				ToTime:   tt.fields.toTime,
				postfix:  tt.fields.postfix,
			}
			if got := t.toTimeString(); got != tt.want {
				t1.Errorf("toTimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

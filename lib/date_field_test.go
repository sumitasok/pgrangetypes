package lib

import (
	"encoding/json"
	"testing"
	"time"
)

const layout = "2006-01-02T15:04:05-07:00"

func TestDateParser_UnmarshalJSON(t *testing.T) {
	inputJson := []byte(`"Mon, 02 Jan 2016 15:04:05 -0700"`)

	type fields struct {
		Time time.Time
	}
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "BasicParse",
			fields: fields{},
			args: args{
				data: inputJson,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &DateParser{
				Time: tt.fields.Time,
			}
			if err := df.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDateParserOneTimeField_UnmarshalJSON(t *testing.T) {
	inputJson := []byte(`{
		"from": "Mon, 02 Jan 2016 15:04:05 -0700"
	}`)

	str := "2016-01-02T15:04:05-07:00"
	timeExample, err := time.Parse(layout, str)
	if err != nil {
		t.Errorf(err.Error())
	}

	type Data struct {
		From DateParser `json:"from"`
	}

	type fields struct {
		data Data
	}
	type args struct {
		data []byte
	}
	type wants struct {
		from time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wants   wants
		wantErr bool
	}{
		{
			name:    "ParseFromStruct",
			fields:  fields{data: Data{}},
			args:    args{data: inputJson},
			wants:   wants{from: timeExample},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &Data{}
			//strings.NewReader(
			if err := json.Unmarshal(tt.args.data, &df); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if df.From.Equal(tt.wants.from) != true {
				t.Errorf("UnmarshalJSON() got = %v, want %v", df.From.String(), tt.wants.from.String())
			}
		})
	}
}

func TestDateParserTwoTimeField_UnmarshalJSON(t *testing.T) {
	inputJson := []byte(`{
		"from": "Mon, 02 Jan 2016 15:04:05 -0700",
		"to": "Mon, 02 Jan 2016 17:04:05 -0700"
	}`)

	strFrom := "2016-01-02T15:04:05-07:00"
	strTo := "2016-01-02T17:04:05-07:00"
	timeFrom, err := time.Parse(layout, strFrom)
	timeTo, err := time.Parse(layout, strTo)
	if err != nil {
		t.Errorf(err.Error())
	}

	type Data struct {
		From DateParser `json:"from"`
		To   DateParser `json:"to"`
	}

	type fields struct {
		data Data
	}
	type args struct {
		data []byte
	}
	type wants struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wants   wants
		wantErr bool
	}{
		{
			name:    "ParseFromStruct",
			fields:  fields{data: Data{}},
			args:    args{data: inputJson},
			wants:   wants{from: timeFrom, to: timeTo},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &Data{}
			//strings.NewReader(
			if err := json.Unmarshal(tt.args.data, &df); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if df.From.Equal(tt.wants.from) != true {
				t.Errorf("From_UnmarshalJSON() got = %v, want %v", df.From.String(), tt.wants.from.String())
			}

			if df.To.Equal(tt.wants.to) != true {
				t.Errorf("To_UnmarshalJSON() got = %v, want %v", df.To.String(), tt.wants.to.String())
			}
		})
	}
}

func TestDateParserTwoTimeFieldInNestedStruct_UnmarshalJSON(t *testing.T) {
	inputJson := []byte(`{
		"daterange": {
			"from": "Mon, 02 Jan 2016 15:04:05 -0700",
			"to": "Mon, 02 Jan 2016 17:04:05 -0700"
		}
	}`)

	strFrom := "2016-01-02T15:04:05-07:00"
	strTo := "2016-01-02T17:04:05-07:00"
	timeFrom, err := time.Parse(layout, strFrom)
	timeTo, err := time.Parse(layout, strTo)
	if err != nil {
		t.Errorf(err.Error())
	}

	type DateRange struct {
		From DateParser `json:"from"`
		To   DateParser `json:"to"`
	}

	type Data struct {
		Daterange DateRange
	}

	type fields struct {
		data Data
	}
	type args struct {
		data []byte
	}
	type wants struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wants   wants
		wantErr bool
	}{
		{
			name:    "ParseFromStruct",
			fields:  fields{data: Data{}},
			args:    args{data: inputJson},
			wants:   wants{from: timeFrom, to: timeTo},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &Data{}
			//strings.NewReader(
			if err := json.Unmarshal(tt.args.data, &df); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if df.Daterange.From.Equal(tt.wants.from) != true {
				t.Errorf("From_UnmarshalJSON() got = %v, want %v", df.Daterange.From.String(), tt.wants.from.String())
			}

			if df.Daterange.To.Equal(tt.wants.to) != true {
				t.Errorf("To_UnmarshalJSON() got = %v, want %v", df.Daterange.To.String(), tt.wants.to.String())
			}
		})
	}
}

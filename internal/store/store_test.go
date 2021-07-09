package store

import (
	"log"
	"reflect"
	"testing"
)

func TestStore_UpdateHistory(t *testing.T) {
	type fields struct {
		data map[string][]Location
		log  *log.Logger
	}
	type args struct {
		orderId  string
		location Location
	}
	tests := map[string]struct {
		fields  fields
		args    args
		want    []Location
		want1   Status
		wantErr bool
	}{
		"updated": {
			fields: fields{data: map[string][]Location{
				"1234": {
					{
						Lat: 1.0,
						Lng: 1.0,
					},
				},
			}},
			args: args{orderId: "1234", location: Location{Lat: 2.0, Lng: 2.0}},
			want: []Location{
				{
					Lat: 2.0,
					Lng: 2.0,
				},
				{
					Lat: 1.0,
					Lng: 1.0,
				},
			},
			want1: Updated,
		},
		"created": {
			fields: fields{data: map[string][]Location{}},
			args:   args{orderId: "1234", location: Location{Lat: 2.0, Lng: 2.0}},
			want: []Location{
				{
					Lat: 2.0,
					Lng: 2.0,
				},
			},
			want1: Created,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := &Store{
				data: tt.fields.data,
				log:  tt.fields.log,
			}
			got, got1, err := s.UpdateHistory(tt.args.orderId, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.UpdateHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.UpdateHistory() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Store.UpdateHistory() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStore_GetHistory(t *testing.T) {
	type fields struct {
		data map[string][]Location
		log  *log.Logger
	}
	type args struct {
		orderId string
		limit   int
	}
	tests := map[string]struct {
		fields  fields
		args    args
		want    []Location
		wantErr bool
	}{
		"ok": {
			fields: fields{data: map[string][]Location{
				"1234": {{Lat: 1.0, Lng: 1.0}},
			}},
			args: args{orderId: "1234"},
			want: []Location{{Lat: 1.0, Lng: 1.0}},
		},
		"err": {
			fields:  fields{data: map[string][]Location{}},
			args:    args{orderId: "1234"},
			wantErr: true,
		},
		"max": {
			fields: fields{data: map[string][]Location{
				"1234": {{Lat: 1.0, Lng: 1.0}, {Lat: 2.0, Lng: 2.0}, {Lat: 3.0, Lng: 3.0}},
			}},
			args: args{orderId: "1234", limit: 1},
			want: []Location{{Lat: 1.0, Lng: 1.0}},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := &Store{
				data: tt.fields.data,
				log:  tt.fields.log,
			}
			got, err := s.GetHistory(tt.args.orderId, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_DeleteHistory(t *testing.T) {
	type fields struct {
		data map[string][]Location
		log  *log.Logger
	}
	type args struct {
		orderId string
	}
	tests := map[string]struct {
		fields  fields
		args    args
		wantErr bool
	}{
		"ok": {
			fields: fields{data: map[string][]Location{"1234": {{Lat: 1.0, Lng: 1.0}}}},
			args:   args{orderId: "1234"},
		},
		"err": {
			fields:  fields{data: map[string][]Location{}},
			args:    args{orderId: "1234"},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := &Store{
				data: tt.fields.data,
				log:  tt.fields.log,
			}
			if err := s.DeleteHistory(tt.args.orderId); (err != nil) != tt.wantErr {
				t.Errorf("Store.DeleteHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

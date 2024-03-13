package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateMetric(t *testing.T) {
	type args struct {
		res http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "zero",
			args: args{
				res: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/update", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateMetric(tt.args.res, tt.args.req)
		})
	}
}

func Test_pushStorage(t *testing.T) {
	type args struct {
		typeMetric string
		nameMetric string
		valMetric  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "zero",
			args: args{
				typeMetric: "counter",
				nameMetric: "count",
				valMetric:  "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pushStorage(tt.args.typeMetric, tt.args.nameMetric, tt.args.valMetric); (err != nil) != tt.wantErr {
				t.Errorf("pushStorage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

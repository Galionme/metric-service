package main

import (
	"testing"
)

func Test_sendServer(t *testing.T) {
	type args struct {
		typeMetric  string
		nameMetric  string
		valueMetric string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "good",
			args: args{
				typeMetric:  "counter",
				nameMetric:  "example",
				valueMetric: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.name = "good"
		})
	}
}

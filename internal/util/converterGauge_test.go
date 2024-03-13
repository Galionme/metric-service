package util

import (
	"testing"
)

func TestGaugeToString(t *testing.T) {
	type args struct {
		gauge interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "good",
			args: args{
				gauge: float64(12.33),
			},
			want:    "12.33",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GaugeToString(tt.args.gauge)
			if (err != nil) != tt.wantErr {
				t.Errorf("GaugeToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GaugeToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToGauge(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "good",
			args: args{
				str: "12.3",
			},
			want:    float64(12.3),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToGauge(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToGauge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToGauge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

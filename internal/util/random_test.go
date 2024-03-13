package util

import "testing"

func TestRandomFloat64(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "good zero",
			args: args{
				min: 0,
				max: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomFloat64(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("RandomFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

package util

import "testing"

func TestCounterToString(t *testing.T) {
	type args struct {
		counter interface{}
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
				counter: int64(120),
			},
			want:    "120",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CounterToString(tt.args.counter)
			if (err != nil) != tt.wantErr {
				t.Errorf("CounterToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CounterToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToCounter(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "good",
			args: args{
				str: "120",
			},
			want:    int64(120),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToCounter(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToCounter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

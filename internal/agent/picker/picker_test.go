package picker

import (
	"reflect"
	"testing"
)

func TestNewPicker(t *testing.T) {
	tests := []struct {
		name string
		want *Picker
	}{
		{
			name: "good",
			want: NewPicker(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPicker(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPicker() = %v, want %v", got, tt.want)
			}
		})
	}
}

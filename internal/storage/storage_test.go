package storage

import (
	"reflect"
	"testing"
)

func TestMemStorage_Get(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue interface{}
		wantOk    bool
	}{
		{
			name: "zero",
			fields: fields{
				data: map[string]interface{}{
					"test": 1,
				},
			},
			args: args{
				key: "test",
			},
			wantValue: 1,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemStorage{
				data: tt.fields.data,
			}
			gotValue, gotOk := m.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMemStorage_Set(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "zero",
			fields: fields{
				data: map[string]interface{}{
					"test": 1,
				},
			},
			args: args{
				key:   "test_key",
				value: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemStorage{
				data: tt.fields.data,
			}
			m.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestNewMemStorage(t *testing.T) {
	tests := []struct {
		name string
		want *MemStorage
	}{
		{
			name: "zero",
			want: &MemStorage{
				data: make(map[string]interface{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

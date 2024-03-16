// Package picker used to simplify work with statistics.
package picker

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	// bigUint represents a constant for a large integer, unit64.
	bigUint = iota
	// mediumUint represents a constant for the average integer, uint32.
	mediumUint
	// bigFloat represents a constant for a large floating point number, float64.
	bigFloat
)

// Picker represents a structure for selecting elements to simplify working with types.
type Picker struct {
	// buf contains a map of string keys and integer values, name and type.
	buf map[string]int
}

// NewPicker creates and returns a new instance of the Picker structure.
func NewPicker() *Picker {
	return &Picker{
		// Initializing the map of the required string keys and integer values, naming and types.
		buf: map[string]int{
			"Alloc":         bigUint,
			"BuckHashSys":   bigUint,
			"Frees":         bigUint,
			"GCSys":         bigUint,
			"HeapAlloc":     bigUint,
			"HeapIdle":      bigUint,
			"HeapInuse":     bigUint,
			"HeapObjects":   bigUint,
			"HeapReleased":  bigUint,
			"HeapSys":       bigUint,
			"LastGC":        bigUint,
			"Lookups":       bigUint,
			"MCacheInuse":   bigUint,
			"MCacheSys":     bigUint,
			"MSpanInuse":    bigUint,
			"MSpanSys":      bigUint,
			"Mallocs":       bigUint,
			"NextGC":        bigUint,
			"OtherSys":      bigUint,
			"PauseTotalNs":  bigUint,
			"StackInuse":    bigUint,
			"StackSys":      bigUint,
			"Sys":           bigUint,
			"TotalAlloc":    bigUint,
			"NumForcedGC":   mediumUint,
			"NumGC":         mediumUint,
			"GCCPUFraction": bigFloat,
		},
	}
}

// GetString simplifies work from meaning and translates.
func (p *Picker) GetString(name string, value reflect.Value) (string, error) {
	typeValue, ok := p.buf[name]
	if !ok {
		return "", fmt.Errorf("value name not found")
	}

	switch typeValue {
	case bigUint:
		return strconv.FormatUint(value.Uint(), 10), nil
	case mediumUint:
		return strconv.FormatUint(value.Uint(), 10), nil
	case bigFloat:
		return strconv.FormatFloat(value.Float(), 'f', -1, 64), nil
	}

	return "", fmt.Errorf("value type not found")
}

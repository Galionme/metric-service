package picker

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	bigUint = iota
	mediumUint
	bigFloat
)

type Picker struct {
	buf map[string]int
}

func NewPicker() *Picker {
	return &Picker{
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

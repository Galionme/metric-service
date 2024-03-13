package util

import (
	"fmt"
	"strconv"
	"strings"
)

func GaugeToString(gauge interface{}) (string, error) {
	if num, ok := gauge.(float64); ok {
		tmp := fmt.Sprintf("%f", num)
		tmp = strings.TrimRight(strings.TrimRight(tmp, "0"), ".")
		return tmp, nil
	}
	return "", fmt.Errorf("is not of type float64")
}

func StringToGauge(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

package util

import (
	"fmt"
	"strconv"
)

func CounterToString(counter interface{}) (string, error) {
	if num, ok := counter.(int64); ok {
		return strconv.FormatInt(num, 10), nil
	}
	return "", fmt.Errorf("is not of type int64")
}

func StringToCounter(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

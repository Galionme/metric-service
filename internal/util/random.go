package util

import (
	"math/rand"
	"time"
)

func RandomFloat64(min, max float64) float64 {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return min + random.Float64()*(max-min)
}

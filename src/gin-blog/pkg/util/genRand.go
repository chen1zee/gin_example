package util

import (
	"math/rand"
	"time"
)

var rand1 = rand.New(rand.NewSource(time.Now().Unix()))

func GenRand() int {
	return rand1.Int()
}

func GenRandMinMax(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

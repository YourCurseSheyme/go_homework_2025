package generators

import (
	"math/rand"
	"time"
)

type IDGenerator func() int

func IncrementalIDGen() IDGenerator {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func RandomIDGen(max int) IDGenerator {
	rand.Seed(time.Now().UnixNano())
	return func() int {
		return rand.Intn(max) + 1
	}
}

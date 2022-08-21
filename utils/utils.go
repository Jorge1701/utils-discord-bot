package utils

import (
	"math/rand"
	"time"
)

func ShuffleList(list []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
}

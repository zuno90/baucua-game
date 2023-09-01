package utils

import (
	"math/rand"
	"time"
)

var dice = []string{"bau", "cua", "tom", "ca", "ga", "nai"}

func RandomResult() []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(dice), func(i, j int) {
		dice[i], dice[j] = dice[j], dice[i]
	})
	return dice[:3]
}

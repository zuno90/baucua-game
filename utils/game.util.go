package utils

import (
	"math/rand"
)

var dice = []string{"bau", "cua", "tom", "ca", "ga", "nai"}

func RandomResult() []string {
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(dice), func(i, j int) {
	// 	dice[i], dice[j] = dice[j], dice[i]
	// })
	// log.Println(dice, 88)
	// return dice[:3]
	result := []string{}
	for i := 1; i <= 3; i++ {
		randomIndex := rand.Intn(len(dice))
		pick := dice[randomIndex]
		result = append(result, pick)
	}
	return result
}

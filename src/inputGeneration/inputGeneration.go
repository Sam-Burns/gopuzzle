package inputGeneration

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("NESW")

func RandomJourney(journeyLength int) string {
	randomChars := make([]rune, journeyLength)
	for index := range randomChars {
		randomChars[index] = letterRunes[rand.Intn(4)]
	}
	return string(randomChars)
}

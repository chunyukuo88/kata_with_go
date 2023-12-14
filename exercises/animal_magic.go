package main

// https://exercism.org/tracks/go/exercises/animal-magic/iterations?idx=1

import (
	"math/rand"
	"time"
)

func RollADie() int {
	number := rand.Intn(20)
	return number + 1
}
func GenerateWandEnergy() float64 {
	rand.Seed(time.Now().UnixNano())
	multiplier := rand.Float64()
	integer := rand.Intn(13)
	return multiplier * float64(integer)
}
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(8, func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}

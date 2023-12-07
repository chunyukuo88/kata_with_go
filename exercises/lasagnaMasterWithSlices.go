// Sauce: https://exercism.org/tracks/go/exercises/lasagna-master/edit

package main

// TODO: define the 'PreparationTime()' function
func PreparationTime(ingredients []string, layers int) int {
	if layers == 0 {
		return 4
	}
	numberOfIngredients := len(ingredients)
	return numberOfIngredients * layers
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (noodles int, sauce float64) {
	for i := range ingredients {
		if ingredients[i] == "noodles" {
			noodles += 50
		}
		if ingredients[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	lastIndex := len(friendsList) - 1
	secretIngredient := friendsList[lastIndex]
	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))

	for i, amount := range quantities {
		scaledQuantities[i] = amount * float64(portions) / 2.0
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

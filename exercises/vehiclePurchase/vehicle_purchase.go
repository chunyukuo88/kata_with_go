package main

// Sauce: https://exercism.org/tracks/go/exercises/vehicle-purchase/edit

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}

}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	sentence := " is clearly the better choice."
	var result string

	if option1 < option2 {
		result = option1 + sentence
		return result
	} else {
		result := option2 + sentence
		return result
	}

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var price float64

	if age < 3 {
		price = originalPrice * .8
		return price
	}
	if age >= 3 && age < 10 {
		price = originalPrice * .7
		return price
	}

	price = originalPrice * .5
	return price
}

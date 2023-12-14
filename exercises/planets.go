package main

// https://exercism.org/tracks/go/exercises/space-age/iterations?idx=1
type Planet string

func isNotPlanet(input Planet) bool {
	planets := []Planet{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	for _, planet := range planets {
		if planet == input {
			return false
		}
	}
	return true
}

func Age(seconds float64, planet Planet) float64 {
	if isNotPlanet(planet) {
		return -1
	}
	earthAge := seconds / 31557600
	var age float64
	switch planet {
	case "Mercury":
		age = earthAge / 0.2408467
	case "Venus":
		age = earthAge / 0.61519726
	case "Mars":
		age = earthAge / 1.8808158
	case "Jupiter":
		age = earthAge / 11.862615
	case "Saturn":
		age = earthAge / 29.447498
	case "Uranus":
		age = earthAge / 84.016846
	case "Neptune":
		age = earthAge / 164.79132
	default:
		age = earthAge
	}
	return age
}

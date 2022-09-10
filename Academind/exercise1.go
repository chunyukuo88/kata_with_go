package main

import (
	"fmt"
)

var pi float32 = 3.14
var circleRadius float32 = 5

func main() {
	var circumference float32 = (2.0 * pi * circleRadius)
	fmt.Printf("For a radius of %v, the circle circumference is %.2f\n", circleRadius, circumference)
}

package main

import (
	"fmt"
	"kata/info"
)

func main() {
	fmt.Println(info.Title)
	weightInput, heightInput := getUserHeightAndWeight()
	weightInt, heightInt := processInputs(weightInput, heightInput)
	bmi := weightInt / (heightInt * heightInt)
	fmt.Printf("Your BMI is %.7f\n", bmi)
}

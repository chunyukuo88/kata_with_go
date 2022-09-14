package main

import (
	"fmt"
	"kata/info"
)

func main() {
	fmt.Println(info.Title)
	weightInput, heightInput := GetUserHeightAndWeight()
	weightInt, heightInt := ProcessInputs(weightInput, heightInput)
	bmi := weightInt / (heightInt * heightInt)
	fmt.Printf("Your BMI is %.7f\n", bmi)
}

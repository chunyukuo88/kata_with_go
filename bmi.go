package main

import (
	"fmt"
	"kata/info"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(info.Title)
	fmt.Println(info.Divider)
	fmt.Println(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Println(info.WeightPrompt)
	heightInput, _ := reader.ReadString('\n')
	strings.Replace(weightInput, "\n", "", -1)
	strings.Replace(heightInput, "\n", "", -1)
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)
	fmt.Print(weight)
	fmt.Print(height)
	bmi := weight / (height * height)
	fmt.Printf("Your BMI is %v", bmi)
}

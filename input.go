package main

import (
	"bufio"
	"fmt"
	"kata/info"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func processInputs(weight string, height string) (float64, float64) {
	weight = strings.Replace(weight, "\n", "", -1)
	height = strings.Replace(weight, "\n", "", -1)
	weightInt, _ := strconv.ParseFloat(weight, 64)
	heightInt, _ := strconv.ParseFloat(height, 64)
	return weightInt, heightInt
}

func getUserHeightAndWeight() (string, string) {
	fmt.Println(info.WeightPrompt)
	weight, _ := reader.ReadString('\n')
	fmt.Println(info.HeightPrompt)
	height, _ := reader.ReadString('\n')
	return weight, height
}

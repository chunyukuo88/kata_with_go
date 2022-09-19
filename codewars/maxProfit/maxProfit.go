/*
https://www.codewars.com/kata/559590633066759614000063/train/go
*/

package maxProfit

func MinMax(input []int) [2]int {
	min := input[0]
	var max int
	for _, s := range input {
		if s >= max {
			max = s
		} else {
			min = s
		}
	}
	return [2]int{min, max}
}

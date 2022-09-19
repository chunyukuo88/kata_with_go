/*
https://www.codewars.com/kata/559590633066759614000063/train/go
*/

package maxProfit

import "sort"

func MinMax(input []int) [2]int {
	sort.Ints(input)
	min := input[0]
	max := input[len(input)-1]
	return [2]int{min, max}
}

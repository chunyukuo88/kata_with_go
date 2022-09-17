/*
https://www.codewars.com/kata/5b077ebdaf15be5c7f000077/train/go
*/

package codewars

import (
	"strconv"
	"strings"
)

func countSheep(num int) string {
	if num == 0 {
		return ""
	}

	var toBeJoined []string

	for i := 1; i <= num; i++ {
		numberPlusSheep := strconv.Itoa(i) + " sheep..."
		toBeJoined = append(toBeJoined, numberPlusSheep)
	}

	return strings.Join(toBeJoined, "")
}

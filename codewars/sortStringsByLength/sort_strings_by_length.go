package sortstringsbylength

// https://www.codewars.com/kata/57ea5b0b75ae11d1e800006c/train/go

import "sort"

func SortByLength(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	return arr
}

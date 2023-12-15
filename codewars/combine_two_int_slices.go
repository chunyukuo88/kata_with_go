package codewars

// https://github.com/chunyukuo88/kata_with_go.git

import "sort"

func includes(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func MergeArrays(arr1, arr2 []int) []int {
	var result []int
	for _, value := range arr1 {
		if !includes(result, value) {
			result = append(result, value)
		}
	}
	for _, value := range arr2 {
		isNewValue := !includes(result, value)
		if isNewValue {
			result = append(result, value)
		}
	}
	sort.Ints(result)
	return result
}

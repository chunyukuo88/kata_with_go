package main

// https://www.codewars.com/kata/56269eb78ad2e4ced1000013/train/go
import "math"

func isNotSquare(number int64) bool {
	root := math.Sqrt(float64(number))
	roundedRoot := math.Floor(root)
	if root-roundedRoot != 0 {
		return true
	}
	return false
}

func FindNextSquare(sq int64) int64 {
	if isNotSquare(sq) {
		return -1
	}
	root := float64(math.Sqrt(float64(sq)))
	return int64(math.Pow((root + 1), 2.0))
}

package maxProfit

import (
	"testing"
)

func TestMinMax_consecutiveIntegers(t *testing.T) {
	arg := []int{1, 2, 3, 4, 5}
	want := [2]int{1, 5}
	got := MinMax(arg)
	if got != want {
		t.Errorf("MinMax(%v) = %v; want %v", arg, got, want)
	}
}

func TestMinMax_largeAndSmall(t *testing.T) {
	arg := []int{2334454, 5}
	want := [2]int{5, 2334454}
	got := MinMax(arg)
	if got != want {
		t.Errorf("MinMax(%v) = %v; want %v", arg, got, want)
	}
}

func TestMinMax_moreNumbers(t *testing.T) {
	arg := []int{3, 100, 4, 1, 5, 1000, 6}
	want := [2]int{1, 1000}
	got := MinMax(arg)
	if got != want {
		t.Errorf("MinMax(%v) = %v; want %v", arg, got, want)
	}
}

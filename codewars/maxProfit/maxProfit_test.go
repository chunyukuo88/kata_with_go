package maxProfit

import (
	"testing"
)

func TestMinMax(t *testing.T) {
	arg := []int{1, 2, 3, 4, 5}
	want := [2]int{1, 5}
	got := MinMax(arg)
	if got != want {
		t.Errorf("MinMax(%v) = %v; want %v", arg, got, want)
	}
}

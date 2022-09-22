package averager

import (
	"testing"
)

// Variadic functions!
func TestAverager_consecutiveIntegers(t *testing.T) {
	var want float32 = 10
	got := Averager(0, 10, 20)
	if got != want {
		t.Errorf("Averager(0, 10, 20) = %v; want %v", got, want)
	}
}

func TestAverager_crunchier_numbers(t *testing.T) {
	var want float32 = 71
	got := Averager(3, 10, 200)
	if got != want {
		t.Errorf("Averager(0, 10, 20) = %v; want %v", got, want)
	}
}

func TestAverager_non_integer_result(t *testing.T) {
	var want float32 = 52.75
	got := Averager(0, 10, 201, 0)
	if got != want {
		t.Errorf("Averager(0, 10, 20) = %v; want %v", got, want)
	}
}

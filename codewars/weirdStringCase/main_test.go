// https://www.codewars.com/kata/52b757663a95b11b3d00062d/train/go
package weirdStringCase

import (
	"testing"
)

func Test_singleWord(t *testing.T) {
	arg := "abc"
	want := "AbC"
	got := ToWeirdCase(arg)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

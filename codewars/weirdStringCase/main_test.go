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

func Test_singleWordFourLettersLong(t *testing.T) {
	arg := "abcd"
	want := "AbCd"
	got := ToWeirdCase(arg)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test_multipleWordsOfVaryingLenth(t *testing.T) {
	arg := "nice tests"
	want := "NiCe TeStS"
	got := ToWeirdCase(arg)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

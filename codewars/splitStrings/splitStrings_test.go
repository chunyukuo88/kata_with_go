// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go
package splitstrings

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Solution_sixLetterString(t *testing.T) {
	arg := "abcdef"
	want := []string{"ab", "cd", "ef"}
	got := Solution(arg)
	if !cmp.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

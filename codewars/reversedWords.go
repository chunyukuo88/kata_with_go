// https://www.codewars.com/kata/51c8991dee245d7ddf00000e/train/go

package codewars

import (
	"strings"
)

func reverse(words []string) []string {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return words
}

func ReverseWords(str string) string {
	split := strings.Split(str, " ")

	return strings.Join(reverse(split), " ")
}

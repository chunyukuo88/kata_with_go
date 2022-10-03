package weirdStringCase

import (
	"strings"
)

func ToWeirdCase(str string) string {
	split := strings.Split(str, " ")
	result := []string{}
	for _, val := range split {
		result = append(result, processSingleWord(val))
	}
	return strings.Join(result, " ")
}

func processSingleWord(str string) string {
	split := strings.Split(str, "")
	result := []string{}
	for i, val := range split {
		if i == 0 || i%2 == 0 {
			result = append(result, strings.ToUpper(val))
		} else {
			result = append(result, val)
		}
	}
	return strings.Join(result, "")
}

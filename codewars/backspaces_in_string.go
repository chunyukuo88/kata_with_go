package codewars

// Source: https://www.codewars.com/kata/5727bb0fe81185ae62000ae3/train/go

func CleanString(s string) string {
	var res []rune

	for _, char := range s {
		if char != '#' {
			res = append(res, char)
		}
		if char == '#' && len(res) > 0 {
			res = res[:len(res)-1]
		}
	}
	return string(res)
}

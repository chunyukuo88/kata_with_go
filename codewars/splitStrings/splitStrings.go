// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go

package splitstrings

func Solution(str string) []string {
	var res []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}

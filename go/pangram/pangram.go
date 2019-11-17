package pangram

import "strings"

const abc = "abcdefghijklmnopqrstunwxyz"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for _, r := range abc {
		if !strings.Contains(input, string(r)) {
			return false
		}
	}

	return true
}

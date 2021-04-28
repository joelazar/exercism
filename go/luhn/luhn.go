package luhn

import (
	"strconv"
	"strings"
)

func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 2 {
		return false
	}

	sum := 0
	mod := len(input) % 2

	for i, char := range input {
		number, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		if i%2 == mod {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		sum += number
	}
	return (sum % 10) == 0
}

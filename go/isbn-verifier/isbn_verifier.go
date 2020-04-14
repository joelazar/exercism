package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(input string) bool {
	sum := 0
	input = strings.Replace(input, "-", "", -1)
	if len(input) != 10 {
		return false
	}
	for i, char := range input {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			if i != len(input)-1 || char != 'X' {
				return false
			}
			num = 10
		}
		sum += num * (10 - i)
	}

	return (sum % 11) == 0
}

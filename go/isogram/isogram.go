package isogram

import "strings"
import "unicode"

// IsIsogram - this function checks if a word is isogram or not
func IsIsogram(input string) bool {
	input = strings.ToLower(input)

	for i, char := range input {
		if !unicode.IsLetter(char) {
			continue
		}
		if strings.Contains(input[i+1:], string(char)) {
			return false
		}
	}
	return true
}

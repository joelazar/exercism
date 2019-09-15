package isogram

import "unicode"

// IsIsogram - this function checks if a word is isogram or not
func IsIsogram(input string) bool {
	var charList []rune
	for _, inputChar := range input {
		if inputChar == '-' || inputChar == ' ' {
			continue
		}
		lowerCasedChar := unicode.ToLower(inputChar)
		for _, char := range charList {
			if lowerCasedChar == char {
				return false
			}
		}
		charList = append(charList, lowerCasedChar)
	}
	return true
}

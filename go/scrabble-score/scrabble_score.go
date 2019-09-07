// This package contains scrabble related functions.
package scrabble

import "unicode"

// Score - this function calculates the score of the given Scrabble word
func Score(word string) int {
	score := 0
	for _, char := range word {
		switch unicode.ToLower(char) {
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		default:
			score++
		}
	}
	return score
}

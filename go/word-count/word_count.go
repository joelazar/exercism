package wordcount

import "strings"
import "unicode"

// Frequency - using this as a return type.
type Frequency map[string]int

var freqMap Frequency

func f(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
}

func filterApostrophe(s string) string {
	if strings.IndexAny(s, "'") == 0 {
		s = s[1:]
	}
	if strings.IndexAny(s, "'") == len(s)-1 {
		s = s[:len(s)-1]
	}
	return s
}

// WordCount - counts how many times is a word can be found in a sentence.
func WordCount(sentence string) Frequency {
	freqMap = make(Frequency)
	words := strings.FieldsFunc(sentence, f)
	for _, word := range words {
		word = filterApostrophe(strings.ToLower(word))
		freqMap[word]++
	}

	return freqMap
}

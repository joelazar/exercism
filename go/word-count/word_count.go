package wordcount

import "strings"
import "unicode"

type Frequency map[string]int // Using this return type.

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

func WordCount(sentence string) Frequency {
	freqMap = make(Frequency)
	words := strings.FieldsFunc(sentence, f)
	for _, word := range words {
		word = filterApostrophe(strings.ToLower(word))
		freqMap[word] += 1
	}

	return freqMap
}

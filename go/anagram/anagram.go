package anagram

import "sort"
import "strings"

// sortString - sorts a string into an alphabetical order
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Detect - this functions detects anagrams of given string from an array of strings and returns with them
func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	subject = strings.ToLower(subject)
	sortedSubject := sortString(subject)

	for _, candidate := range candidates {
		if subject != strings.ToLower(candidate) && sortedSubject == sortString(strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

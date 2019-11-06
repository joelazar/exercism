package isogram

import "sort"
import "strings"
import "unicode"

// IsIsogram - this function checks if a word is isogram or not
func IsIsogram(input string) bool {
	runeList := []rune(strings.ToLower(input))
	sort.Slice(runeList, func(i, j int) bool {
		return runeList[i] < runeList[j]
	})

	for i := 0; i < len(runeList)-1; i++ {
		if !unicode.IsLetter(runeList[i]) {
			continue
		}
		if runeList[i] == runeList[i+1] {
			return false
		}
	}
	return true
}

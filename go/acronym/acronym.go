package acronym

import "strings"

// Abbreviate - abbreviate
func Abbreviate(s string) string {

	var acronym string
	capitalWords := strings.Fields(strings.ToUpper(strings.Replace(s, "-", " ", -1)))

	for _, r := range capitalWords {
		acronym += string(r[0])
	}

	return acronym
}

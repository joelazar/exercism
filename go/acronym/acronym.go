package acronym

import "strings"

func Abbreviate(s string) string {

	var acronym string
	bigwords_withouthyphen := strings.Fields(strings.ToUpper(strings.Replace(s, "-", " ", -1)))

	for _, r := range bigwords_withouthyphen {
		acronym += string(r[0])
	}

	return acronym
}

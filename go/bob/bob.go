package bob

import "strings"
import "unicode"

func anyLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// Hey - Whatever :)
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	} else if anyLetter(remark) && strings.ToUpper(remark) == remark {
		if remark[len(remark)-1:] == "?" {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	} else if remark[len(remark)-1:] == "?" {
		return "Sure."
	} else {
		return "Whatever."
	}
}

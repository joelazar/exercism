package twofer

import "fmt"

const TEXT string = "One for %s, one for me."

func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf(TEXT, name)
}

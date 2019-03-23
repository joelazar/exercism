package twofer

import "fmt"

const text string = "One for %s, one for me."

// One for X, one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf(text, name)
}

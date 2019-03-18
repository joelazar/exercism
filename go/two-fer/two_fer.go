package twofer

import "fmt"

const text string = "One for %s, one for me."

// ShareWith - One for X, one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf(TEXT, name)
}

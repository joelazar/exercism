package proverb

import "fmt"

const (
	baseSentence = "For want of a %s the %s was lost."
	lastSentence = "And all for the want of a %s."
)

// Proverb - a function which creates the new proverb of the given input string list
func Proverb(rhyme []string) []string {
	proverb := []string{}
	for i := range rhyme {
		if i == len(rhyme)-1 {
			proverb = append(proverb, fmt.Sprintf(lastSentence, rhyme[0]))
		} else {
			proverb = append(proverb, fmt.Sprintf(baseSentence, rhyme[i], rhyme[i+1]))
		}
	}
	return proverb
}

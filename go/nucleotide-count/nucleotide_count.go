package dna

import "errors"

type Histogram map[rune]int

type DNA string

func validNucleotide(c rune) bool {
	if c != 'C' && c != 'G' && c != 'A' && c != 'T' {
		return false
	}
	return true
}

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, value := range d {
		if validNucleotide(value) {
			h[value]++
		} else {
			return h, errors.New("not valid nucleotide found")
		}
	}
	return h, nil
}

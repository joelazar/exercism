package strand

import "strings"

func rnaMap(input rune) rune {
	switch input {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		return 'Z'
	}
}

// ToRNA - converting dna string to rna one
func ToRNA(dna string) string {
	return strings.Map(rnaMap, dna)
}

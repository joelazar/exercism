package protein

import "errors"

// ErrStop - This error indicates a terminating codon.
var ErrStop error = errors.New("stop")

// ErrInvalidBase - This error indicates an invalid codon in the sequence.
var ErrInvalidBase error = errors.New("invalid base")

// FromCodon - converts given codon to an amino acid.
func FromCodon(input string) (string, error) {
	switch input {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA - converts given RNA sequences into a protein.
func FromRNA(input string) ([]string, error) {
	rna := []string{}
	for i := 0; i < len(input)-2; i += 3 {
		codon, err := FromCodon(input[i : i+3])
		if err != nil {
			if err == ErrInvalidBase {
				return rna, err
			}
			break
		}
		rna = append(rna, codon)
	}
	return rna, nil
}

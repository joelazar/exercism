package hamming

import "errors"

// Distance - Hamming distance measuring between two DNS strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Two strings are not same length")
	}

	var counter int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			counter++
		}
	}

	return counter, nil
}

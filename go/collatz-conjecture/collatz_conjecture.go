package collatzconjecture

import "errors"

// CollatzConjecture - calculate the steps of the conjecture
func CollatzConjecture(input int) (int, error) {
	var counter int
	if input <= 0 {
		return 0, errors.New("input is not valid")
	}
	for input != 1 {
		counter++
		if input%2 == 0 {
			input /= 2
		} else {
			input = 3*input + 1
		}
	}
	return counter, nil
}

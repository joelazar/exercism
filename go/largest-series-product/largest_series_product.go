package lsproduct

import (
	"fmt"
	"strconv"
)

func NProduct(input []int, span int) int {
	prod := 1
	for i := 0; i < span; i++ {
		prod *= input[i]
	}
	return prod
}

func LargestSeriesProduct(input string, span int) (int, error) {
	numinput := []int{}
	for _, char := range input {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return 0, fmt.Errorf("array contains non-digit: %v", err)
		}
		numinput = append(numinput, num)
	}

	if span > len(numinput) {
		return 0, fmt.Errorf("span must be smaller than the array")
	}
	if span < 0 {
		return 0, fmt.Errorf("span must be non negative")
	}

	maxprod := NProduct(numinput, span)
	for i := 1; i < len(numinput)-(span-1); i++ {
		newprod := NProduct(numinput[i:], span)
		if maxprod < newprod {
			maxprod = newprod
		}
	}

	return maxprod, nil
}

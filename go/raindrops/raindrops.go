package raindrops

import "strconv"

// Convert - converts a number to strings of pling/plang/plong
func Convert(input int) string {
	var output string

	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(input)
	}

	return output
}

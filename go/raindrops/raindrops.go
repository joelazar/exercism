package raindrops

import "strconv"

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

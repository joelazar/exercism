package encode

import "strconv"
import "unicode"

var counter = 1
var currRune rune
var encodedString = ""

func appendData() {
	if counter == 1 {
		encodedString = encodedString + string(currRune)
	} else {
		encodedString = encodedString + strconv.Itoa(counter) + string(currRune)
	}
	counter = 1
}

// RunLengthEncode - Encodes string into RLE form.
func RunLengthEncode(input string) string {
	encodedString = ""
	counter = 1
	for i, r := range input {
		if i == 0 {
			currRune = r
			continue
		}
		if r == currRune {
			counter++
		} else {
			appendData()
			currRune = r
		}
		if i == len(input)-1 {
			appendData()
		}
	}

	return encodedString
}

func findLastNumIndex(input string) int {
	for i, r := range input {
		if !unicode.IsNumber(r) {
			return i
		}
	}

	return len(input) - 1
}

// RunLengthDecode - Decodes string from RLE form.
func RunLengthDecode(input string) string {
	decodedString := ""
	for len(input) != 0 {
		lastIndex := findLastNumIndex(input)
		if num, err := strconv.Atoi(string(input[0:lastIndex])); err == nil {
			for i := 0; i < num; i++ {
				decodedString = decodedString + string(input[lastIndex])
			}
			input = input[lastIndex+1:]
		} else {
			decodedString = decodedString + string(input[0])
			input = input[1:]
		}
	}

	return decodedString
}

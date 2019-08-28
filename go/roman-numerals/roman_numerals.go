package romannumerals

import "errors"

var romanNumberConverter = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var romanNumbers = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

// ToRomanNumeral - convert arabic numbers to roman one
func ToRomanNumeral(arabicNum int) (string, error) {
	romanNumber := ""
	if arabicNum <= 0 || arabicNum > 3000 {
		return romanNumber, errors.New("Incorrect input received")
	}
	for arabicNum > 0 {
		for _, num := range romanNumbers {
			if arabicNum >= num {
				romanNumber += romanNumberConverter[num]
				arabicNum -= num
				break
			}
		}
	}
	return romanNumber, nil
}

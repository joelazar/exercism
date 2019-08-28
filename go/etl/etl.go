package etl

import "strings"

// Transform - this function transforms the old legacydata structure to a new format
func Transform(legacyData map[int][]string) map[string]int {
	newMap := map[string]int{}
	for key, value := range legacyData {
		for _, char := range value {
			char = strings.ToLower(char)
			newMap[char] = key
		}
	}

	return newMap
}

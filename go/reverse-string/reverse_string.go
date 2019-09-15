package reverse

// Reverse - reversing the given string
func Reverse(input string) string {
	reversed := []rune{}

	for _, value := range input {
		reversed = append([]rune{value}, reversed...)
	}
	return string(reversed)
}

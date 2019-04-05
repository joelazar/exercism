package accumulate

func Accumulate(inputList []string, fn func(string) string) []string {
	for i, value := range inputList {
		inputList[i] = fn(value)
	}
	return inputList
}

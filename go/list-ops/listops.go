package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (list IntList) Length() int {
	var counter int
	for range list {
		counter++
	}
	return counter
}

func (list IntList) Append(otherList IntList) IntList {
	merged := make(IntList, list.Length()+otherList.Length())
	for i, value := range list {
		merged[i] = value
	}
	for i, value := range otherList {
		merged[list.Length()+i] = value
	}
	return merged
}

func (list IntList) Map(fn unaryFunc) IntList {
	for i, value := range list {
		list[i] = fn(value)
	}
	return list
}

func (list IntList) Foldl(fn binFunc, input int) int {
	for _, value := range list {
		input = fn(input, value)
	}
	return input
}

func (list IntList) Foldr(fn binFunc, input int) int {
	for _, value := range list.Reverse() {
		input = fn(value, input)
	}
	return input
}

func (list IntList) Reverse() IntList {
	var temp int
	for i := 0; i < list.Length()/2; i++ {
		temp = list[i]
		list[i] = list[list.Length()-i-1]
		list[list.Length()-i-1] = temp
	}
	return list
}

func (list IntList) Filter(fn predFunc) IntList {
	var lastIndex int
	for _, value := range list {
		if fn(value) {
			list[lastIndex] = value
			lastIndex++
		}
	}
	return list[0:lastIndex]
}

func (list IntList) Concat(otherLists []IntList) IntList {
	for _, otherList := range otherLists {
		list = list.Append(otherList)
	}
	return list
}

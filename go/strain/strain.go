package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(fn func(int) bool) Ints {
	var filtered = Ints(nil)
	for i, value := range ints {
		if fn(value) {
			filtered = append(filtered, ints[i])
		}
	}
	return filtered
}

func (ints Ints) Discard(fn func(int) bool) Ints {
	var filtered = Ints(nil)
	for i, value := range ints {
		if !fn(value) {
			filtered = append(filtered, ints[i])
		}
	}
	return filtered
}

func (lists Lists) Keep(fn func([]int) bool) Lists {
	var filtered = Lists(nil)
	for i, value := range lists {
		if fn(value) {
			filtered = append(filtered, lists[i])
		}
	}
	return filtered
}

func (strings Strings) Keep(fn func(string) bool) Strings {
	var filtered = Strings(nil)
	for i, value := range strings {
		if fn(value) {
			filtered = append(filtered, strings[i])
		}
	}
	return filtered
}

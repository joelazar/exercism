package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0

	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			if i%divisor == 0 {
				sum += i
				break
			}
		}
	}

	return sum
}

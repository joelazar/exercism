package prime

var primes []int

func IsPrime(num int) bool {
	for _, prime := range primes {
		if prime == 0 {
			return true
		}
		if num%prime == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	counter := 1
	number := 2
	primes = make([]int, n)
	primes[0] = 2

	if n == 1 {
		return primes[0], true
	}

	for {
		if IsPrime(number) {
			primes[counter] = number
			counter += 1
			if counter == n {
				return number, true
			}
		}
		number += 1
	}
}

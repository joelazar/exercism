package pythagorean

type Triplet [3]int

// Integer power: compute a**b using binary powering algorithm
// See Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	list := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if Pow(a, 2)+Pow(b, 2) == Pow(c, 2) {
					list = append(list, Triplet{a, b, c})
				}
			}
		}
	}
	return list
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	list := []Triplet{}
	for a := 1; a <= p/3; a++ {
		for b := a; b <= p/2; b++ {
			for c := b; c <= p/2; c++ {
				if a+b+c != p {
					continue
				}
				if Pow(a, 2)+Pow(b, 2) == Pow(c, 2) {
					list = append(list, Triplet{a, b, c})
				}
			}
		}
	}
	return list
}

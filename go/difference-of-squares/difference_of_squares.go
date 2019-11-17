package diffsquares

import "math"

// SquareOfSum - calculate the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	return int(math.Pow(float64((n*(n+1))/2), 2))
}

// SumOfSquares - calculate the sum of the first squared n natural numbers
func SumOfSquares(n int) int {
	return ((n * (n + 1) * (2*n + 1)) / 6)
}

// Difference - calculate the difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)

}

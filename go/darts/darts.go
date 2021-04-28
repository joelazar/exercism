package darts

import "math"

const (
	outerCircleRadius  = 10
	middleCircleRadius = 5
	innerCircleRadius  = 1
)

// Score - calculate the scores by the given dart coordinates
func Score(x, y float64) int {
	squareroot := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	if squareroot > outerCircleRadius {
		return 0
	} else if squareroot > middleCircleRadius {
		return 1
	} else if squareroot > innerCircleRadius {
		return 5
	} else {
		return 10
	}
}

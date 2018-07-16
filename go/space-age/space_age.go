package space

import "math"

type Planet string

var planets = map[Planet]float64{
	"Earth":   31557600,
	"Mercury": 0.2408467 * 31557600,
	"Venus":   0.61519726 * 31557600,
	"Mars":    1.8808158 * 31557600,
	"Jupiter": 11.862615 * 31557600,
	"Saturn":  29.447498 * 31557600,
	"Uranus":  84.016846 * 31557600,
	"Neptune": 164.79132 * 31557600,
}

func Age(seconds float64, planet Planet) float64 {
	return math.Round((seconds/planets[planet])/0.01) * 0.01
}

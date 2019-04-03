package scale

import "fmt"
import "container/ring"
import "strings"

var scaleWithSharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var scaleWithFlats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var useSharps = []string{"C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}

func stringContains(a string, list []string) bool {
	for _, value := range list {
		if a == value {
			return true
		}
	}
	return false
}

func ringContains(a string, ringlist *ring.Ring) *ring.Ring {
	for i := 0; i < ringlist.Len(); i++ {
		if ringlist.Value == strings.Title(a) {
			break
		}
		ringlist = ringlist.Next()
	}
	return ringlist
}

func fillRing(a *ring.Ring, list []string) {
	for _, value := range list {
		a.Value = value
		a = a.Next()
	}
}

func Scale(tonic string, interval string) []string {
	var (
		workingRing = ring.New(12)
		outputScale []string
		useSharps   = stringContains(tonic, useSharps)
	)

	if useSharps {
		fillRing(workingRing, scaleWithSharps)
	} else {
		fillRing(workingRing, scaleWithFlats)
	}

	workingRing = ringContains(tonic, workingRing)

	if interval == "" {
		for i := 0; i < workingRing.Len(); i++ {
			outputScale = append(outputScale, fmt.Sprintf("%v", workingRing.Value))
			workingRing = workingRing.Next()
		}
	} else {
		for _, char := range interval {
			outputScale = append(outputScale, fmt.Sprintf("%v", workingRing.Value))
			switch string(char) {
			case "m":
				workingRing = workingRing.Next()
			case "M":
				workingRing = workingRing.Next().Next()
			case "A":
				workingRing = workingRing.Next().Next().Next()
			}
		}
	}

	return outputScale
}

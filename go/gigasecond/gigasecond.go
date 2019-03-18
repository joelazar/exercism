package gigasecond

import "math"
import "time"

// AddGigasecond - adding gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow10(9)) * time.Second)
}

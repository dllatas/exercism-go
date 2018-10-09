// Package gigasecond provides a function
// to perform addition on a time variables
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes as an input a time and returns
// the sum of that time and 1 000 000 000 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}

// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond provides a function to add a gigasecond (1,000,000,000 seconds) to a given time.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds a gigasecond (1,000,000,000 seconds) to the given time.
func AddGigasecond(t time.Time) time.Time {
	// Define a gigasecond as 1,000,000,000 seconds
	const gigasecond = 1e9 // 1,000,000,000 seconds

	// Add the gigasecond to the given time
	t = t.Add(gigasecond * time.Second)
	return t
}

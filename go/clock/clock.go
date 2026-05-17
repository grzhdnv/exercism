package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {

	day := 60 * 24
	total := (h*60 + m) % day

	for total < 0 {
		total += day
	}

	h = total / 60
	m = total % 60

	// for m < 0 {
	// 	m += 60
	// 	h -= 1
	// }

	// for h < 0 {
	// 	h += 24
	// }

	// if m >= 60 {
	// 	h += m / 60
	// 	m = m % 60
	// }
	// if h >= 24 {
	// 	h = h % 24
	// }

	return Clock{
		hours:   h,
		minutes: m,
	}
}

func (c Clock) Add(m int) Clock {
	totalMinutes := c.hours*60 + c.minutes
	hours := totalMinutes / 60
	minutes := totalMinutes % 60
	return New(hours, minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

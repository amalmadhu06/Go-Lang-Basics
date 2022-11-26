//ParseDuration parses a duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

package main

import (
	"fmt"
	"time"
)

func main() {
	// h for hour
	// m for minute
	// s for second
	// u for nanosecond

	// x, err := time.ParseDuration("10h")
	// x, err := time.ParseDuration("5h10m10s")
	x, err := time.ParseDuration("8µs")
	fmt.Printf("%d", x.Nanoseconds())
	fmt.Println()
	fmt.Println(err)

}

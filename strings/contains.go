package main

import (
	"fmt"
	"strings"
)

// Contains reports whether substr is within s.

func main() {
	fmt.Println(strings.Contains("amal", "a")) // --> true
	fmt.Println(strings.Contains("amal", "s")) // --> false
}

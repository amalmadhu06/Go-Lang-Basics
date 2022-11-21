package main

import (
	"fmt"
	"strings"
)

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code
// points in s.

func main() {

	fmt.Println(strings.Count("amal", "a"))   // 2  (it contains two a)
	fmt.Println(strings.Count("amal", ""))    // 5  (empty space after each letter is counted)
	fmt.Println(strings.Count("amal", "str")) // 0
	fmt.Println(strings.Count("amal", "am"))  // 1
}

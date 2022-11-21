package main

import (
	"fmt"
	"strings"
)

//Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b.

func main() {
	fmt.Println(strings.Compare("One", "One"))       // --> 0
	fmt.Println(strings.Compare("One", "one"))       // --> -1
	fmt.Println(strings.Compare("One", "two"))       // --> -1
	fmt.Println(strings.Compare("One", "something")) // --> -1
	fmt.Println(strings.Compare("some", "One"))      // --> 1
	fmt.Println(strings.Compare(" ", "One"))         // --> -1

}

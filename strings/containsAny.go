package main

import (
	"fmt"
	"strings"
)

func main() {
	// ContainsAny reports whether any Unicode code points in
	// chars are within s.

	fmt.Println(strings.ContainsAny("amal", "pqr")) // --> false
	fmt.Println(strings.ContainsAny("amal", "ram")) // --> false

}

package main

import (
	"fmt"
	"strings"
)

// ContainsRune reports whether the Unicode code point r is within s.

func main() {
	fmt.Println(strings.ContainsRune("amal", 97)) // --> true (because unicode of "a" is 97)
	fmt.Println(strings.ContainsRune("dude", 98)) // --> false (because unicode
	// 98 means "b". It is not present in the string.)

}

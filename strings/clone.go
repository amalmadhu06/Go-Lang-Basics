package main

import (
	"fmt"
	"strings"
)

func main() {

	// clone makes a fresh copy of the string to a new variable

	s1 := "sample string"

	s2 := strings.Clone(s1)

	fmt.Println(s1) // --> sample string
	fmt.Println(s2) // --> sample string
}

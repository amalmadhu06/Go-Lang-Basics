package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	args := []string{"amal", "madhu", "ranni"}
	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}
}

func Concat(parts ...string) (string, error) { //variadic func with strings as input, string and error as return value
	if len(parts) == 0 {
		return "", errors.New("No strings supplied") //creating a new error and returning it
	}
	return strings.Join(parts, " "), nil
}

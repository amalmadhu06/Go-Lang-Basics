package main

import (
	"errors"
	"fmt"
	"strings"
)

func Concat(parts ...string) (string, error) { //variadic func with strings as input, string and error as return value
	if len(parts) == 0 {
		return "", errors.New("No strings supplied") //creating a new error and returning it
	}
	return strings.Join(parts, " "), nil
}

func main() {
	fmt.Println(Concat())                //output : No strings supplied
	fmt.Println(Concat("Amal", "Madhu")) //output : Amal Madhu ,nil
}

package main

import (
	"fmt"
	"strings"
)

//EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal
//  under simple Unicode case-folding, which is a more general form of
//  case-insensitivity.

func main() {

	fmt.Println(strings.EqualFold("amal", "Amal")) //true
	fmt.Println(strings.EqualFold("go", "Go"))     //true
	fmt.Println(strings.EqualFold("go", "Gooo"))   //fale

}

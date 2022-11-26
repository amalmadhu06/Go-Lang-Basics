//string literals are string constants obtained from concatinating a
//sequence of characters. They are of two types.
// Raw string literals		--> written inside back quotes `amal`
// Interpreted string literals --> written inside quotes "amal"

package main

import "fmt"

func main() {

	x := `amal` // --> Raw string literals
	y := "amal" // --> Interpreted string literals

	fmt.Println(x, y)
}

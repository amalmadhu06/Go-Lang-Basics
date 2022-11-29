package main

import "fmt"

func main() {

	//The 'make' allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make’s return type is the same as the type of its argument, not a pointer to it.
	s := make([]int, 5)
	fmt.Println(s) // [0,0,0,0,0]

	// The new built-in function allocates memory. The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.
	a := new([]int)
	b := append(*a, 1)

	

	fmt.Println(b)
}

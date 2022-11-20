package main

import "fmt"

// func main() {
// 	x := func(x int) int {
// 		fmt.Println("x is returned")
// 		return x * 1
// 	}
// }

func main() {
	returnFunc("hello")()
}

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

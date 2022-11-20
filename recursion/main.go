package main

import "fmt"

func main() {

	fmt.Println(factorial(20))
}

func factorial(x uint64) uint64 { //program to find the factotial of a number
	if x <= 1 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

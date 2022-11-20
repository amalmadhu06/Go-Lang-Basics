package main

import "fmt"

func main() {
	fmt.Println(fib(25))

}

func fib(x int) int {
	if x <= 1 {
		return x
	} else {
		return fib(x-1) + fib(x-2)
	}

}

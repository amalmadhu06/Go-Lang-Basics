package main

import (
	"fmt"
)

func main() {

	var input int
	fmt.Println("Please enter a number")
	fmt.Scan(&input)

	//checking for prime
	var flag int = 0
	for i := 1; i <= input; i++ {
		if (input % i) == 0 {
			flag++
		}
	}
	if flag == 2 {
		fmt.Println("Number is prime")
	} else {
		fmt.Println("Number is not prime")
	}
}

package main

import "fmt"

func main() {

	var input int
	fmt.Println("Please enter the value")
	fmt.Scan(&input)

	for i := 1; i <= 10; i++ {
		fmt.Println(i, " x ", input, " = ", input*i)
	}
}

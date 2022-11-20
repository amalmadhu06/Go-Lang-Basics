package main

import "fmt"

func main() {

	var limit int
	fmt.Println("Please enter the limit")
	fmt.Scan(&limit)

	fmt.Println("limit is", limit)
	sum := 0
	for i := 1; i <= limit; {
		sum = sum + i
		i = i + 2

	}

	fmt.Println("Sum of the first", limit, " odd numbers is : ", sum)
}

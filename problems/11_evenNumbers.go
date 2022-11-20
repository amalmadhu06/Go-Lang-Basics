//count the even numbers in an array

package main

import "fmt"

func main() {

	var arr [100]int
	var limit int

	fmt.Println("Please enter the limit of the array")
	fmt.Scan(&limit)

	fmt.Println("Enter the values of the array")

	for i := 0; i < limit; i++ {
		fmt.Scan(&arr[i])
	}

	count := 0
	for _, c := range arr {
		if c%2 == 1 {
			count++
		}
	}

	fmt.Println("Number of odd numbers is :", count)

}

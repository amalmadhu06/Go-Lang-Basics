package main

import (
	"fmt"
)

func main() {

	var arr [100]int
	var limit int

	fmt.Println("Enter the size of the array")
	fmt.Scan(&limit)

	fmt.Println("Enter the values of the array")

	for i := 0; i < limit; i++ {
		fmt.Scan(&arr[i])
	}

	//printing the array

	fmt.Println()
	fmt.Println("Array is : ")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d  ", arr[i])
	}

	temp := 0

	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			if arr[j] > arr[i] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println()
	fmt.Println("Array after sorting is : ")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d  ", arr[i])
	}

}

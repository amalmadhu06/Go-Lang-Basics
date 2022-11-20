package main

import "fmt"

func main() {

	var size int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&size)

	var array = make([]int, size)
	var multiArr = make([]int, size-1)

	fmt.Println("Enter the values of the array")

	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	// asssining values to new array

	for i := 0; i < size-1; i++ {
		multiArr[i] = array[i] * array[i+1]
	}
	fmt.Println(multiArr)
}

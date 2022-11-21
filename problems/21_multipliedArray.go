package main

import "fmt"

func main() {

	var size int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&size)

	var array = make([]int, size)
	// var array []int

	// var array [100]int
	var multiArr = make([]int, size-1)
	// var multiArr [100]int

	fmt.Println("Enter the values of the array")

	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	// asssining values to new array

	for i := 0; i < size-1; i++ {
		multiArr[i] = array[i] * array[i+1]
	}
	// fmt.Println(multiArr)

	//printing the array
	for i := 0; i < size-1; i++ {
		fmt.Printf("%v  ,", multiArr[i])
	}

	// 1, 2, 3, 4
	// 2, 6, 12

}

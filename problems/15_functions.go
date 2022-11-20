package main

import "fmt"

func main() {
	var a1 [100]int
	var s1 int
	a1, s1 = getArrayy()

	displayArray(a1, s1)

}

func getArrayy() ([100]int, int) {
	var size int
	var arr [100]int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&size)

	fmt.Println("Enter the values of the array")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	return arr, size
}

func displayArray(arr [100]int, size int) {

	for i := 0; i < size; i++ {
		fmt.Printf("%d  ", arr[i])
	}
}

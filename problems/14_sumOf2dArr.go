package main

import "fmt"

func main() {
	var arr1 [100][100]int
	var arr2 [100][100]int
	var sum [100][100]int
	var size int

	fmt.Println("Please enter the size of the arrays")
	fmt.Scan(&size)

	fmt.Println("Please enter the values of the first array")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&arr1[i][j])
		}
	}

	fmt.Println("Please enter the values of the second array")

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}

	fmt.Println("First array is : ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d  ", arr1[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Second array is : ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d  ", arr2[i][j])
		}
		fmt.Println()
	}

	//calculating the sum

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	// printing the sum array

	fmt.Println("Sum array is : ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d  ", sum[i][j])
		}
		fmt.Println()
	}
}

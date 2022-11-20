package main

import "fmt"

func main() {

	displayArray(addArray(getArray()))

}

//function to receive two slices

func getArray() ([][]int, [][]int, int) {

	var size int
	var sliceRow []int
	var slice1 [][]int
	var slice2 [][]int

	fmt.Println("Please enter the size of the array")
	fmt.Scan(&size)

	fmt.Println("Please enter the values of the first array")

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var element int
			fmt.Scan(&element)
			sliceRow = append(sliceRow, element)
		}

		slice1 = append(slice1, sliceRow)
		sliceRow = nil

	}

	fmt.Println("Please enter the values of the second array")

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var element int
			fmt.Scan(&element)
			sliceRow = append(sliceRow, element)
		}
		slice2 = append(slice2, sliceRow)
		sliceRow = nil

	}

	return slice1, slice2, size
}

// function to add two slices

func addArray(s1, s2 [][]int, s int) [][]int {
	size := s
	var slice3 []int
	var sumSlice [][]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			slice3 = append(slice3, s1[i][j]+s2[i][j])
		}
		sumSlice = append(sumSlice, slice3)
		slice3 = nil
	}
	return sumSlice
}

//function to print all three slices

func displayArray(s1 [][]int) {
	fmt.Println("Sum array is : ", s1)
}

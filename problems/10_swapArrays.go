package main

import "fmt"

func main() {

	var array1 [100]int
	var array2 [100]int
	var limit, temp int

	fmt.Println("Enter the size of the arrays")
	fmt.Scan(&limit)

	fmt.Println("Enter the values of the first array")
	for i := 0; i < limit; i++ {
		fmt.Scan(&array1[i])
	}

	fmt.Println("Enter the values of the second array")
	for i := 0; i < limit; i++ {
		fmt.Scan(&array2[i])
	}

	fmt.Println("\nFirst array before swapping : ")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d  ", array1[i])
	}

	fmt.Println("\n Second array before swapping : ")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d  ", array2[i])
	}

	//swapping the values

	for i := 0; i < limit; i++ {
		temp = array1[i]
		array1[i] = array2[i]
		array2[i] = temp
	}

	//arrays after swapping

	fmt.Println("\nFirst array after swapping : ")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d  ", array1[i])
	}

	fmt.Println("\n Second array after swapping : ")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d  ", array2[i])
	}

}

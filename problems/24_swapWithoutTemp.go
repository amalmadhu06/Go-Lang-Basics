package main

import "fmt"

func main() {

	var x int
	var y int

	fmt.Println("Please enter x value")
	fmt.Scan(&x)
	fmt.Println("Please enter y value")
	fmt.Scan(&y)

	//swapping without temp
	x = x + y
	y = x - y
	x = x - y

	fmt.Println("x after swapping is : ", x)
	fmt.Println("y after swapping is : ", y)
}

package main

import "fmt"

func main() {
	var num1 int
	var num2 float64
	var sum float64
	fmt.Println("Enter the integer : ")
	fmt.Scan(&num1)
	fmt.Println("Enter the decimal value : ")
	fmt.Scan(&num2)

	sum = float64(num1) + num2
	
	fmt.Println("Sum of the entered number is :", sum)
}

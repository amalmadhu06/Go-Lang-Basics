package main

import "fmt"

func main() {
	var num1, num2 float64
	var operation int
	fmt.Println("Enter the first number :")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number :")
	fmt.Scan(&num2)

	fmt.Println("Select 1 for : Addition \nSelect 2 for : Subtraction \nSelect 3 for : multiplication \nSelect 4 for : division")
	fmt.Scan(&operation)

	switch operation {
	case 1:
		addition(num1, num2)
	case 2:
		subtraction(num1, num2)
	case 3:
		multiplication(num1, num2)
	case 4:
		division(num1, num2)
	default:
		fmt.Println("Wrong input... Please enter a value from 1 to 4")
	}

}

func addition(number1 float64, number2 float64) float64 {

	result := number1 + number2
	fmt.Println(result)
	return result
}

func subtraction(number1 float64, number2 float64) float64 {
	result := number1 - number2
	fmt.Println(result)
	return result
}

func multiplication(number1 float64, number2 float64) float64 {
	result := number1 * number2
	fmt.Println(result)
	return result
}

func division(number1 float64, number2 float64) float64 {

	result := number1 / number2
	fmt.Println(result)
	return result
}

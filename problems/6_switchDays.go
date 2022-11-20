package main

import "fmt"

func main() {

	var input int
	fmt.Println("Please enter the input number from 1 to 7")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Please enter a value from 1 to 7")
	}

}

package main

import "fmt"

//  * and & operator
//  new

func main() {

	x := 10

	fmt.Println("Initial value of x is :", x)
	fmt.Println("Address if x is : ", &x)

	/*
		name of the variable : x
		value stored in x 	 : 10
		address of x 		 : 0xc00000e098
	*/

	xPointer := &x
	fmt.Println("We are storing the address of x to xPointer, and xPointer is now : ", xPointer)

	*xPointer = 11

	fmt.Println("Value of x after changing value using pointer is : ", x)
	fmt.Println("*xPointer : ", *xPointer)
	fmt.Println("xPointer : ", xPointer)

}

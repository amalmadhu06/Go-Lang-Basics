package main

import "fmt"

//for changing value in  another function we need to use pointer
// here we are using pointer in the function squareVal to make 
// to the variable in the function main.


func main() {
	x := 10
	squareVal(&x)
}

func squareVal(x *int) {
	*x = *x * *x
	fmt.Println(*x)
}

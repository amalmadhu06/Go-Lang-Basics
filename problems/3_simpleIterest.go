package main

import "fmt"

func main() {
	var p int
	var r float64
	var n float64
	var si float64

	fmt.Println("Enter the principal amount")
	fmt.Scan(&p)
	fmt.Println("Enter the interest rate")
	fmt.Scan(&r)
	fmt.Println("Enter the number of years")
	fmt.Scan(&n)

	si = (float64(p) * r * n)/100

	fmt.Println("Simple interest is : ", si)

}
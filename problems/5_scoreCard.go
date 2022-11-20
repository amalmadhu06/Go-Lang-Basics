package main

import "fmt"

func main() {
	var totalMark float64

	fmt.Println("Please enter your total mark :")
	fmt.Scan(&totalMark)

	if totalMark >= 90 && totalMark <=100 {
		fmt.Println("A")
	} else if totalMark >= 80 && totalMark < 90 {
		fmt.Println("B")
	} else if totalMark >= 70 && totalMark < 80 {
		fmt.Println("C")
	} else if totalMark >= 60 && totalMark < 70 {
		fmt.Println("D")
	} else if totalMark >= 50 && totalMark < 60 {
		fmt.Println("E")
	} else if totalMark < 50 {
		fmt.Println("Failed")
	} else {
		fmt.Println("Please enter  correct value")
	}
}

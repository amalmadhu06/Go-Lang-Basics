package main

import "fmt"

func main() {

	var mark float64

	fmt.Println("Please enter your score: ")
	fmt.Scan(&mark)

	if mark < 50 {
		fmt.Println("Sorry, you have failed in the exam")
	} else {
		fmt.Println("You have cleared the exam ")
	}

}


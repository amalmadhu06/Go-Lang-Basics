package main

import "fmt"

func main() {
	var written, lab, assignment float64

	fmt.Println("Enter your score for written test")
	fmt.Scan(&written)
	fmt.Println("Enter your score for lab exam")
	fmt.Scan(&lab)
	fmt.Println("Enter your score for assignment")
	fmt.Scan(&assignment)

	//calculating weighted average

	score := (written*70)/100 + (lab*20)/100 + (assignment*10)/100
	fmt.Println("Your score is : ", score)
}

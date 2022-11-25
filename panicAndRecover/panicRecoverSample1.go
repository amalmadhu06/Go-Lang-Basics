package main

import "fmt"

func main() {
	defer tryToRecover()
	fmt.Println("Main function started")
	panicCreator()
}

func panicCreator() {
	fmt.Println("Panic creator started")
	panic("created a panic")
}

func tryToRecover() {
	if err := recover(); err != nil {
		fmt.Println("recovered from error")
	}
}

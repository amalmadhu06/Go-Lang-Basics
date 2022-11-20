package main

import "fmt"

func main() {

	test1()
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func test1() {
	fmt.Println("one")
	fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("four")
	fmt.Println("five")
}

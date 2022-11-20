package main

import "fmt"

func main() {
	test := func(x int) int {
		return x * -1
	}

	test2(test)

}

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

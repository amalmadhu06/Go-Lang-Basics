package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Please enter the string")
	fmt.Scan(&input)
	input = strings.ToLower(input)

	flag := 0
	for i := 0; i < len(input)/2; i++ {

		if input[i] != input[len(input)-1-i] {
			flag = 1
			fmt.Println("String is not palindrome")
			break
		}
	}

	
	if flag == 0 {
		fmt.Println("String is palindrome")
	}


}

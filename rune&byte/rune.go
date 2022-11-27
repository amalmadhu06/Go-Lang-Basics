package main

import "fmt"

func main() {

	str := "GOLANG"
	slice1 := []rune(str)

	var result []int

	for i := 0; i < len(slice1); i++ {
		result = append(result, int(slice1[i]))
	}

	fmt.Println(result)

}

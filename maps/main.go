package main

import "fmt"

func main() {
	score := make(map[string]int)

	score["Amal"] = 10
	score["Yadhu"] = 11
	score["Shameem"] = 12

	// fmt.Println(score)

	for i, j := range score {
		fmt.Printf("Score of %v is %v \n", i, j)
	}
	fmt.Println(score)

}

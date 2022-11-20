package main

import "fmt"

func main() {

	var s = []int{1, 2, 3}
	fmt.Println(len(s))
	fmt.Println(s)
	var sNew = append(s, 4, 5, 6)
	fmt.Println(sNew)

	sOne := sNew[2:4]
	fmt.Println(sOne)

	sTwo := sNew[:4]
	fmt.Println(sTwo)

	sThree := sNew[:]
	fmt.Println(sThree)
}

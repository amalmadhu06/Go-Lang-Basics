package main

import "fmt"

func main() {
	num := 1
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if j <= i {

				fmt.Printf(" %d ", num)
				num++
			}
		}
		fmt.Println()
	}
}

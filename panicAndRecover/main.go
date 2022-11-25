package main

import "fmt"

func main() {
	for i := 0; i <= 11; i++ {
		if i == 7 {
			defer func() {

				if r := recover(); r != nil {
					fmt.Println("recovered in f", r)
				}
			}()
			panic("Oops... its 9")
		} else {
			fmt.Println(i)
		}
	}
}

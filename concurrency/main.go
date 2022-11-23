package main

import (
	"fmt"
	"time"
)

func main() {

	go printerFunc("first --1")
	time.Sleep(time.Second * 1)
	go printerFunc("second ----2")
	time.Sleep(time.Second * 1)
	printerFunc("third ------3")

	time.Sleep(time.Second * 5)

}

func printerFunc(s string) {

	for i := -0; true; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 1000)
	}
}

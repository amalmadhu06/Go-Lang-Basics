package main

import (
	"fmt"
	"time"
)

func main() {

	helloPrinter("hi")
	go helloPrinter("hello")

}

func helloPrinter(s string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}

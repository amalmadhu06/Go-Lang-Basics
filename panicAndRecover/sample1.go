package main

import (
	"fmt"
	"time"
)

func main() {
	defer recoverFunc()
	fmt.Println("started")
	fmt.Println("sleeping for...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
	panic("sleeping over")
	fmt.Println("after panic")

}

func recoverFunc() {
	if err := recover(); err != nil {
		fmt.Println("Trapped in panic")
	}
}

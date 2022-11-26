package main

import (
	"fmt"
	"time"
)

func sleepyFunc() {
	time.Sleep(time.Millisecond * 700)
}

func main() {
	//calculating time took to execute the function
	t1 := time.Now()
	sleepyFunc()
	t2 := time.Now()
	timeTook := t2.Sub(t1)
	fmt.Println(timeTook)
}

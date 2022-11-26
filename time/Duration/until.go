package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	time.Sleep(time.Millisecond * 500)
	calculatedTime := time.Until(startTime)
	fmt.Println(calculatedTime)
}

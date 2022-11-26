package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	time.Sleep(time.Millisecond * 700)
	elapsed := time.Since(startTime)
	fmt.Println(elapsed)

}

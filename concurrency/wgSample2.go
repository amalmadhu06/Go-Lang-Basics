package main

import (
	"fmt"
	"sync"
	"time"
)

//time elapsed without go routines	: 1.9354598s
//time elapsed with go routines		: 703.1356ms

func main() {
	startTime := time.Now()
	hugeTask()
	fmt.Println("Time elapsed is : ", time.Since(startTime))
}

func hugeTask() {

	var wg sync.WaitGroup
	
	wg.Add(4)
	go func() {
		fmt.Println("task 1")
		time.Sleep(time.Millisecond * 500)
		wg.Done()
	}()

	go func() {
		fmt.Println("task 2")
		time.Sleep(time.Millisecond * 700)
		wg.Done()
	}()

	go func() {
		fmt.Println("task 3")
		time.Sleep(time.Millisecond * 300)
		wg.Done()
	}()

	go func() {
		fmt.Println("task 4")
		time.Sleep(time.Millisecond * 400)
		wg.Done()
	}()

	wg.Wait()

}

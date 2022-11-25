package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	startTime := time.Now()
	wg.Add(4)
	multiTask(&wg)
	wg.Wait()
	fmt.Println("Time taken is : ", time.Since(startTime))
}

func multiTask(wg *sync.WaitGroup) {
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
}

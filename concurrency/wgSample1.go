package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(4)
	now := time.Now()
	go func() {
		defer wg.Done()
		fmt.Println("one")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("two")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("three")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("four")
	}()

	wg.Wait()
	fmt.Println("elapsed time is : ", time.Since(now))
	fmt.Println("main function is done")
}

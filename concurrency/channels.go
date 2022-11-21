package main

import (
	"fmt"
	"sync"
)

func main() {

	myChan := make(chan int, 3)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		// fmt.Println(<-myChan)
		val, openOrNot := <-myChan
		fmt.Println(val, openOrNot)
		wg.Done()
	}(myChan, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// myChan <- 6
		myChan <- 5
		// close(myChan)
		wg.Done()
	}(myChan, wg)

	wg.Wait()
}

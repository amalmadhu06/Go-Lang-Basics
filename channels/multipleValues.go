package main

import (
	"fmt"
	"time"
)

func multiValueReceiver(channel1 chan int) { //function with receives values from channels multiple times
	value := 0
	for value >= 0 {
		value = <-channel1
		time.Sleep(time.Millisecond * 500)
		fmt.Println(value)
	}
}

func main() {

	sampleChannel := make(chan int)           //making a channel of type int
	go multiValueReceiver(sampleChannel)      //calling the function
	s := []int{1, 2, 3, 0, 5, 6, 7, 8, 9, -1} //storing values in a slice
	for _, value := range s {                 //sending them over channel
		sampleChannel <- value
	}

	time.Sleep(time.Millisecond * 500)        //for pausing main before ending
	fmt.Println("Main function completed   ") //
}

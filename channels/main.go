package main

import (
	"fmt"
	"time"
)

func printValue(c chan int) { //function which receives values through channel
	num := <-c                             // receiving value through channel
	fmt.Println(num)                       //printing the received value
	fmt.Println("                       ") //

}

func main() {
	channel1 := make(chan int)             // making a channel with type int
	go printValue(channel1)                //calling the function which receives value from the channel
	channel1 <- 10                         //passing value to the channel
	time.Sleep(time.Millisecond * 300)     //setting timeout
	fmt.Println("Main function completed") //
}

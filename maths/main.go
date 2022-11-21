package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//trying to generate a random number using math.rand
	//using math.rand will always returns the same value here
	//so we need to use rand.see and pass time to it to generate random seed

	rand.Seed(int64(time.Now().UnixNano()))
	fmt.Println(rand.Intn(5))
}

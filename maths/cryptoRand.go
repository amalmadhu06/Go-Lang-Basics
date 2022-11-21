package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// generating random number using crypto

	randomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNum)

}

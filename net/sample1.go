package main

import (
	"fmt"
	"net"
)

func main() {

	ipAddress, err := net.LookupIP("google.com")

	fmt.Println(ipAddress)
	fmt.Println(err)
}

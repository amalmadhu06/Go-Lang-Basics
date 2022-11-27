package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, _ := net.LookupIP("facebook.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	iprecords2, err := net.LookupIP("tiktok.com")
	for _, ip := range iprecords2 {
		fmt.Println(ip)
	}
	fmt.Println(err)
}

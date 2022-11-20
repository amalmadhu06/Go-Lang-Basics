package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	endpointList := []string{
		"https://google.com",
		"https://facebook.com",
		"https://go.dev",
		"https://github.com",
		"https://amazon.com",
	}

	for _, web := range endpointList {
		go getEndpoint(web)
		wg.Add(1)
	}

	wg.Wait()
}

func getEndpoint(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops...Something is wrong with", endpoint)
	} else {
		fmt.Println("status code for ", endpoint, " : ", res.StatusCode)

	}
}

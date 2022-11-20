package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

// creating a variable to use mutex
var mut sync.Mutex

var signals []string

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
	fmt.Println(signals)
}

func getEndpoint(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops...Something is wrong with", endpoint)
	} else {
		// locking the memory access to signals with mutex
		mut.Lock()
		signals = append(signals, endpoint)
		//unloacing memory access to signals with mutex
		mut.Unlock()

		fmt.Println("status code for ", endpoint, " : ", res.StatusCode)

	}
}

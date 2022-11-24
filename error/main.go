package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println(passChecker(99))
}

func fileFetcher(s string) {
	file, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(file))
	}
}

func passChecker(score float64) (bool, error) {
	status := false
	var err error
	if score < 50 {
		status = false
	} else if score > 100 {
		err = errors.New("enter a valid score")
	} else {
		status = true
	}

	return status, err
}

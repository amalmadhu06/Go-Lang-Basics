package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.com/") //making http GET Request
	fmt.Println(res.Status, err)                //printing respose status and error
	body, _ := ioutil.ReadAll(res.Body)         //reading the body from response
	fmt.Println(string(body))                   //
	res.Body.Close()                            //closes the connection
}

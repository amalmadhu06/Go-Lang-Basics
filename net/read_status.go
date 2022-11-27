package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "google.com: 80")      //connects over tcp
	fmt.Println("Error :", err)                               //
	fmt.Println()                                             //
	fmt.Fprintf(connection, "GET / HTTP/1.0\r\n\r\n")         //sends string over the connection
	status, _ := bufio.NewReader(connection).ReadString('\n') //first response line
	fmt.Println(status)                                       //

}

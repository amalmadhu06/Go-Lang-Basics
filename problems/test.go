package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("test file")
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	CleanupMessage(message)
}

func CleanupMessage(oldMsg string) string {

	res := strings.ReplaceAll(oldMsg, "*", "")
	res = strings.TrimSpace(res)
	fmt.Println(res)
	return res
}

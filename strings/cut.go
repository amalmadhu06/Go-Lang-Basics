package main

import (
	"fmt"
	"strings"
)

//Cut slices s around the first instance of sep, returning the text before
//  and after sep. The found result reports whether sep appears in s. If sep
//  does not appear in s, cut returns s, "", false

func main() {

	before, after, found := strings.Cut("amal", "a")
	fmt.Print(before, " , ", after, " , ", found)

	before1, after1, found1 := strings.Cut("strings", "a")
	fmt.Print(before1, " , ", after1, " , ", found1)

	before2, after2, found2 := strings.Cut("strangs", "a")
	fmt.Print(before2, " , ", after2, " , ", found2)

}

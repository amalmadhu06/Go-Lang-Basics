package main

import "fmt"

func main() {
	a := "മ"   //length --> 3
	b := "ല"   //length --> 3
	c := "യ"   //length --> 3
	c1 := "യാ" //length --> 6
	d := "ള"   //length --> 3
	d1 := "ളം" //length --> 3
	x := "a"   //length --> 1

	fmt.Printf("%s", x)

	fmt.Println(len(x))
	fmt.Println(len(a))

	fmt.Println(len(b))
	fmt.Println(len(c))
	fmt.Println(len(c1))
	fmt.Println(len(d))
	fmt.Println(len(d1))

}

package main

import "fmt"

type Car struct {
	make  string
	year  int
	color string
}

func main() {

	var city = Car{
		make:  "honda",
		year:  2020,
		color: "grey",
	}

	city1 := Car{}
	city1.make = "honda"
	city1.color = "white"
	city1.year = 2009

	fmt.Println(city.color)
	fmt.Println(city1)
}

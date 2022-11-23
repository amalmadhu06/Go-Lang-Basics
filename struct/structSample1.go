package main

import "fmt"

func main() {

	amal := details{
		name:  "Amal",
		age:   22,
		place: "Ranni",
	}

	fmt.Println(amal)
}

type details struct {
	name  string
	age   int
	place string
}

func (s details) detailsPrinter() {
	fmt.Println(s.age, s.name, s.place)
}

type address struct {
	place    string
	district string
}

func (s address) detailsPrinter() {
	fmt.Println(s.district, s.place)
}

type printer interface {
	detailsPrinter()
}

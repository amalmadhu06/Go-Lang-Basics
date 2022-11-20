package main

import "fmt"

//creating a struct
//writing a function to create new objects with the struct
//this function will return a pointer to the object created

func main() {
	// fmt.Println(*initPerson())
	fmt.Printf(" main: %p \n", initPerson())

}

type person struct {
	name string
	age  int
}

func initPerson() *person {

	m := person{
		name: "Amal",
		age:  50}

	fmt.Printf("initPerson : %p \n", &m)
	return &m

}

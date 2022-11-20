package main

import "fmt"

func main() {

	// Amal := User{
	// 	Name:   "Amal",
	// 	Age:    22,
	// 	Email:  "amalmadhu@gmail.com",
	// 	Status: true,
	// }

	var amal User = User{
		Name:  "Amal",
		Age:   22,
		Email: "amalmadhu@gmail.com",
	}
	// fmt.Printf("%+v\n", Amal)
	fmt.Printf("Name is : %v \nAge is : %v ", amal.Name, amal.Age)

}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

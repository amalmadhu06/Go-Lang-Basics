package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}

// Welcome greets a person by name.
func Welcome(name string) string {
	message1 := "Welcome to my party, " + name + "!"
	return message1

}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	message2 := "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
	return message2
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	message3 := "Welcome to my party, " + name + "!\nYou have been assigned to table 0" + strconv.Itoa(table) + ". Your table is on the " + direction + ", exactly " + fmt.Sprintf("%0.1fm", distance) + " from here.\nYou will be sitting next to " + neighbor
	return message3
}

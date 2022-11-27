package main

import "fmt"

type Shape struct {
	name string
	size int
}

func NewShape(name string) Shape {
	return Shape{
		name: name,
		size: 100,
	}
}
func main() {

	s := NewShape("rectangle")
	fmt.Println(s.name, s.size)
}

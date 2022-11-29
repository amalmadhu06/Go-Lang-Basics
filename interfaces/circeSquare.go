package main

import "fmt"

type Calculator interface {
	area() float64 //method set (method signature)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	a := 3.14 * c.radius * c.radius
	return a
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) area() float64 {
	a := r.length * r.breadth
	return a
}

func areaPrinter(s Calculator) {
	fmt.Printf("Area of the figure is  %v ", s.area())
}

func main() {
	slice1 := []Calculator{
		Rectangle{length: 10, breadth: 5},
		Circle{radius: 5},
		Rectangle{length: 12, breadth: 6},
	}

	for _, v := range slice1 {
		areaPrinter(v)
		fmt.Println()
	}
}

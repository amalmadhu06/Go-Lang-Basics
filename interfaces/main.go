package main

import (
	"fmt"
	"math"
)

//shape interface

type shape interface {
	area() float64    //method sets
	circumf() float64 //method sets	
}

type square struct { //creating a struct
	length float64
}

type circle struct { //creating a struct
	radius float64
}

//struct method for square

func (s square) area() float64 {

	return s.length * s.length
}

func (s square) circumf() float64 {

	return s.length * 4
}

// struct methods for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

//function which uses shape

func printShapeInfo(s shape) {
	fmt.Printf("Area of the shape : %T , Circumference of the shape : %T", s, s.area())
}

// main function
func main() {

	shapes := []shape{
		square{length: 10},
		circle{radius: 20},
		circle{radius: 15},
		square{12},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println()
	}
}

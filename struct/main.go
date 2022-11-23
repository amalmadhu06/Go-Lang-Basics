package main

import "fmt"

func main() {

	rectangle1 := rect{
		width:  10,
		height: 5}

	// fmt.Println("area: ", rectangle1.area())
	// fmt.Println("perim:", rectangle1.perim())

	rp := &rectangle1 //pointer to the rectangle1
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

type rect struct {
	width, height int
}

func (r *rect) area() int { //this is a struct method
	return r.width * r.height
}

func (r rect) perim() int { //this is a struct method
	return 2*r.width + 2*r.height
}

package main

import "fmt"

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	s := Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
	return s
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	s := Track{
		distance: distance,
	}
	return s
}

func Drive(car Car) Car {

	// distanceDriven := car.batteryDrain
	s := Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      100 - car.batteryDrain,
		distance:     car.distance}
	return s
}

// => Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}

func main() {
	// speed := 5
	// batteryDrain := 2
	// car := NewCar(speed, batteryDrain)
	// fmt.Println(car)
	// distance := 800
	// fmt.Println(NewTrack(distance))
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Println(Drive(car))

}

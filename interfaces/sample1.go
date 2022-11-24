package main

import "fmt"

func main() {
	t1 := train{
		speed:       200,
		distance:    2000,
		cost:        500,
		boarding:    "TVM",
		destination: "Goa",
	}
	t2 := train{
		speed:       150,
		distance:    6000,
		cost:        500,
		boarding:    "TVM",
		destination: "Gujarat",
	}

	b1 := bus{
		speed:    60,
		distance: 300,
		cost:     300,
	}

	b2 := bus{
		speed:    50,
		distance: 200,
		cost:     150,
	}

	details := []travelCalc{
		t1,
		t2,
		b1,
		b2,
	}

	for _, x := range details {
		detailsPrinter(x)

	}

}

func detailsPrinter(m travelCalc) {
	fmt.Println("______________________________________")
	fmt.Println("Estimated time is : ", m.avgTimeCal())
	fmt.Println("Cost per km is  : ", m.costPerKmCalc())
	fmt.Println("______________________________________")
}

type travelCalc interface {
	avgTimeCal() float64
	costPerKmCalc() float64
}

type train struct {
	speed       float64
	distance    float64
	cost        int
	boarding    string
	destination string
}

func (t train) avgTimeCal() float64 {
	avgSpeed := t.distance / t.speed
	return avgSpeed
}

func (t train) costPerKmCalc() float64 {
	costPerKm := t.distance / float64(t.cost)
	return costPerKm
}

type bus struct {
	speed    float64
	distance float64
	cost     int
}

func (b bus) avgTimeCal() float64 {
	avgSpeed := b.distance / b.speed
	return avgSpeed
}

func (b bus) costPerKmCalc() float64 {
	costPerKm := b.distance / float64(b.cost)
	return costPerKm
}

package main

import "fmt"

type Vehicule struct {
	vehiculeBrand string
	model         string
	topSpeed      int
}

type Speed interface {
	howFast() string

}

func (v Vehicule) getTopSpeed() int {
	return v.topSpeed
}

// doesn't change the state of our vehicule
func (v Vehicule) setTopSpeed(newSpeed int) {
	v.topSpeed = newSpeed
}

// doeas change the state of our vehicule
// func (v *Vehicule) setTopSpeed2(newSpeed int) {
// 	v.topSpeed = newSpeed
// }

func (v Vehicule) howFast() string {
	return "normal speed"
}

type Car struct {
	Vehicule
}

func (c Car) howFast() string {
	return "somehow fast speed"
}

type Truck struct {
	Vehicule
}

func (t Truck) howFast() string {
	return "too slow in terms of speed"
}

func printSpeed(s Speed) {
	fmt.Println(s.howFast())
}

func main() {

	volvo := Truck{
		Vehicule: Vehicule{
			vehiculeBrand: "volvo",
			model:         "ef34",
			topSpeed:      550,
		},
	}

	audi := Car{
		Vehicule: Vehicule{
			vehiculeBrand: "Audi",
			model:         "rs5",
			topSpeed:      320,
		},
	}

	fmt.Println(audi)

	fmt.Println("old speed: ", volvo.getTopSpeed())
	volvo.setTopSpeed(220)
	fmt.Println("new speed: ", volvo.getTopSpeed())

	printSpeed(volvo)
	printSpeed(audi)

}

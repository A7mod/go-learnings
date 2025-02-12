package interfacee

import "fmt"

type Vehicle interface {
	Start()
	Stop()
	getspeed() int
}

type Car struct {
	speed int
}

// Implementing Vehicle interface for Car
func (c Car) Start() {
	fmt.Println("Car is Starting...")
}

func (c Car) Stop() {
	fmt.Println("Car is Stopping...")
}

func (c Car) getspeed() int {
	return c.speed
}

type Bike struct {
	speed int
}

func (b Bike) Start() {
	fmt.Println("Bike is Starting...")
}

func (b Bike) Stop() {
	fmt.Println("Bike is Stopping...")
}

func (b Bike) getspeed() int {
	return b.speed
}

func TestVehicle(v Vehicle) {
	v.Start()
	fmt.Println("Current Speed:", v.getspeed(), "km/h.")
	v.Stop()
}

func ShowVehicle() {
	car := Car{speed: 120}
	bike := Bike{speed: 90}

	fmt.Println("Testing Car:")
	TestVehicle(car)

	fmt.Println("Testing Bike:")
	TestVehicle(bike)

	fmt.Println("----------------------------")
}

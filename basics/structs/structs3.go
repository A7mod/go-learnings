package structs

import "fmt"

type Car struct {
	Brand string
	Year  int
}

// UpdateYear method has a pointer receiver (*Car), allows modification of the actual struct instance
func (c *Car) UpdateYear(newYear int) {
	c.Year = newYear // modifies the original struct instance
}

func ShowCar() {
	// creating an instance of Car struct
	car := Car{Brand: "Toyota", Year: 2021}
	fmt.Println("Before Update Car:", car)

	// calling UpdateYear().
	car.UpdateYear(2024)
	fmt.Println("After Update:", car)

	fmt.Println("----------------------------")

}

package structs

import "fmt"

type Person struct {
	Name string
	Age  int
}

func StructBasics() {
	p := Person{"Amit", 25} // this is a struct instance
	fmt.Println("This is the struct:", p)

	// Accessing items:
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("----------------------------")

}

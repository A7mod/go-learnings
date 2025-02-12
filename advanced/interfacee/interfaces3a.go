package interfacee

import "fmt"

// In this example, we see how Methods are used with Interfaces.

type Speaker interface {
	Speak()
}

type Walker interface {
	Walk()
}

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Println(h.Name, "says: Hello there!")
}

func (h Human) Walk() {
	fmt.Println(h.Name, "Is Walking")
}

func PerformAction(i interface{}) {
	switch v := i.(type) { // this is type switching
	case Speaker:
		fmt.Println("This is a Speaker")
		v.Speak()

	case Walker:
		fmt.Println("This is a Walker")
		v.Walk()

	default:
		fmt.Println("Unknown Type hai ye bhai!")
	}
}

func ShowReality() {
	// Creating Human Instance
	person := Human{Name: "Kesha", Age: 32}

	// Calling function with different interfaces
	PerformAction(person)
}

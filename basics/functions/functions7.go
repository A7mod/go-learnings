package functions

import "fmt"

// defining struct named Person
type Person struct {
	Name string
	Age  int
}

// this is a method -> (Greet()) associated with Person - the struct
// (p Person) is a receiver, i.e. it belongs to Person.

func (p Person) Greet() { // dekho zyda mujhe bhi smjh nhi aaya abhi
	fmt.Println("Hi!, My name is", p.Name, "My name is", p.Name, "My age is:", p.Age, "Slim Shady")
}

func ShowMethods() {
	// new 'instance' of Person
	person1 := Person{Name: "Mike", Age: 32}
	// person1 is calling the Greet()
	person1.Greet()
	fmt.Println("----------------------------")
}

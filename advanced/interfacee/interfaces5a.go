package interfacee

import "fmt"

type Animals interface {
	Speaka() string
}

type Pet interface {
	Animals
	LoveOwner() string
}

type Doggy struct{}

func (d Doggy) Speaka() string {
	return "Woof mofo Woof!"
}

func (d Doggy) LoveOwner() string {
	return "Wags tail vigorously!"
}

func Showpets() {
	var myPet Pet = Doggy{}
	fmt.Println("Dog says:", myPet.Speaka()) // Output: Dog says: Woof!
	fmt.Println("Dog reacts:", myPet.LoveOwner())

	fmt.Println("----------------------------")
}

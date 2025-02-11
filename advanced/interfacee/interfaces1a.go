package interfacee

import "fmt"

// Interface : Animal

type Animal interface {
	Speak() string
}

type Dog struct{}

type Cat struct{}

func (d Dog) Speak() string {
	return "Woof Woof!"
}

func (c Cat) Speak() string {
	return "Meowww!"
}

func ShowSpeak() {
	animals := []Animal{Dog{}, Cat{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	fmt.Println("----------------------------")
}

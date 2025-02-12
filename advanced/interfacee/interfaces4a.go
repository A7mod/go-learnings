// Polymorphism example

package interfacee

import "fmt"

type Speakers interface {
	Speak() string
	GetName() string
}

type Doggo struct{}
type Caty struct{}
type Humane struct{}

func (d Doggo) Speak() string {
	return "Woof Woof!"
}

func (c Caty) Speak() string {
	return "Meowwwww!"
}

func (h Humane) Speak() string {
	return "Hello, Hi"
}

func (d Doggo) GetName() string {
	return "Dog"
}
func (c Caty) GetName() string {
	return "Cat"
}
func (h Humane) GetName() string {
	return "Human"
}

func MakeSpeak(s Speakers) {
	fmt.Println(s.GetName(), "says:", s.Speak())
}

func ShowSpeaker() {
	dog := Doggo{}
	cat := Caty{}
	human := Humane{}

	MakeSpeak(dog)
	MakeSpeak(cat)
	MakeSpeak(human)

	fmt.Println("----------------------------")

}

package structs

import "fmt"

type Rectangle struct {
	length, width float64
}

// Method associated with Struct
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func ShowRectangle() {
	rectangle := Rectangle{length: 10, width: 4}
	fmt.Println("Dimensions:", rectangle)
	fmt.Println("Area of Rectangle:", rectangle.Area())

	fmt.Println("----------------------------")
}

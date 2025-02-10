package functions

import "fmt"

// An example of an anonymous function
func ShowAnonfunc() {
	add := func(a, b int) int { // gumnaami baba pro max
		return a + b
	}
	fmt.Println("Anonymous Function Sum:", add(3, 4))

	fmt.Println("----------------------------")
}

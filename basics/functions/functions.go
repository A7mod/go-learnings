package functions

import "fmt"

// function with parameters
func greet(name string) {
	fmt.Println("Hello ba'ys", name)
}

// function with return value
func add(a int, b int) int {
	return a + b
}

func ShowFunctions() {
	greet("Pitter patter")

	sum := add(7, 9)
	fmt.Println("Sum:", sum)

	fmt.Println("----------------------------")

}

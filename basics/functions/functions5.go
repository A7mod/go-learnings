package functions

import "fmt"

func MulByTwo(x int) int {
	return x * 2
}

func ProcessNumber(n int, fn func(int) int) int {
	return fn(n)
}

func ShowFuncAsArgument() {
	fmt.Println("Function as Argument:")
	result := ProcessNumber(5, MulByTwo)
	fmt.Println("Processed Result:", result)
	fmt.Println("----------------------------")
}

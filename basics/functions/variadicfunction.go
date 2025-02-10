package functions

import "fmt"

func Sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func ShowVariadicFunc() {
	fmt.Println("Variadic Function......")
	fmt.Println("Sum:", Sum(10, 20, 30, 40, 50))

	fmt.Println("----------------------------")
}

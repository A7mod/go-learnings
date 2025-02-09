package basics

import "fmt"

func AddMul(a int, b int) (int, int) {
	sum := a + b
	product := a * b

	return sum, product
}

// isme toh bada chod h bhai.
func Represent() {
	n1, n2 := 4, 5
	sum, product := AddMul(n1, n2)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)

}

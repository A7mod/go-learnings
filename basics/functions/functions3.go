package functions

import "fmt"

func Divide(a, b float64) (result float64, err string) {
	if b == 0 {
		err = "Can't divide by zero. I suggest you let that one marinate!"
		return
	}
	result = a / b
	return
}

// this is also one more type of function. More like a way to write a function
func NamedReturns() {
	res, err := Divide(10, 2)
	if err != "" {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
	fmt.Println("----------------------------")
}

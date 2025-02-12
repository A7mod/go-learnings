package interfacee

import "fmt"

func PrintValue(val interface{}) {
	fmt.Println("Value", val)
}

func ShowValues() {
	PrintValue(42)
	PrintValue("Hello")
	PrintValue(99.98)

	fmt.Println("----------------------------")
}

package interfacee

import "fmt"

func GetData(input int) interface{} {
	if input == 1 {
		return 42
	} else if input == 2 {
		return "Hello, Pitter Patter"
	} else {
		return 98.99
	}
}

func ShowData() {
	data1 := GetData(1)
	data2 := GetData(2)
	data3 := GetData(5)

	// Using type assertion to extract actual values
	switch v := data1.(type) {
	case int:
		fmt.Println("Received an integer:", v)
	case string:
		fmt.Println("Received a string:", v)
	case float64:
		fmt.Println("Received a float:", v)
	default:
		fmt.Println("Unknown type!")
	}

	// Printing remaining values
	fmt.Println("Data2:", data2)
	fmt.Println("Data3:", data3)

	fmt.Println("----------------------------")

}

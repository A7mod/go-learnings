package basics

import "fmt"

func CheckCondition() {
	age := 55

	if age >= 26 {
		fmt.Println("Some criteria!")
	} else {
		fmt.Println("Other criteria.")
	}
	fmt.Println("----------------------------")

}

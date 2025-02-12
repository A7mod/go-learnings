package interfacee

import "fmt"

func ShowYosulf() {
	profile := map[string]interface{}{
		"name":   "Miracles",
		"age":    23,
		"skills": []string{"Go", "AWS", "Docker", "K8s"},
		"active": true,
	}

	// Accessing values dynamically
	fmt.Println("Name:", profile["name"])
	fmt.Println("Age:", profile["age"])
	fmt.Println("Skills:", profile["skills"])
	fmt.Println("Active:", profile["active"])

	fmt.Println("----------------------------")
}

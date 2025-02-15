package main

import (
	"fmt"

	"github.com/A7mod/go-learnings/advanced/concurrency"
	"github.com/A7mod/go-learnings/advanced/interfacee"
)

func main() {

	fmt.Println("\n Running Advanced ... ")

	// advanced as in another main folder-ch
	fmt.Println("\n--- Interfaces ---")
	interfacee.ShowSpeak()
	interfacee.ShowAccounts()
	interfacee.ShowValues()
	interfacee.ShowData()
	interfacee.ShowYosulf()
	interfacee.ShowReality()
	interfacee.ShowVehicle()
	interfacee.ShowSpeaker()
	interfacee.ShowPayments()
	interfacee.Showpets()

	fmt.Println("\n--- Concurrency ---")
	//concurrency.ShowGoroutines()
	//concurrency.ShowGoroutines2()
	//concurrency.ShowBasicChannel()
	concurrency.ShowChannelFunction()

}

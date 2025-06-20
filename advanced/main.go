package main

import (
	"fmt"

	"github.com/A7mod/go-learnings/advanced/concurrency/con2"
)

func main() {

	fmt.Println("\n Running Advanced ... ")

	// advanced as in another main folder-ch
	fmt.Println("\n--- Interfaces ---")
	// interfacee.ShowSpeak()
	// interfacee.ShowAccounts()
	// interfacee.ShowValues()
	// interfacee.ShowData()
	// interfacee.ShowYosulf()
	// interfacee.ShowReality()
	// interfacee.ShowVehicle()
	// interfacee.ShowSpeaker()
	// interfacee.ShowPayments()
	// interfacee.Showpets()

	fmt.Println("\n--- Concurrency ---")
	//concurrency.ShowGoroutines()
	//concurrency.ShowGoroutines2()
	//concurrency.ShowBasicChannel()
	//concurrency.ShowChannelFunction()
	//concurrency.ShowKaise()
	//concurrency.ShowChannelBak()
	//concurrency.ShowBufferedChannelBlocking()
	//concurrency.ShowBufferedLoop()

	fmt.Println("\n--- Concurrency 2 = Con2 ---")

	// con2.ShowChannelDirection()
	// con2.ShowSelectBasics()
	// con2.ShowSelectNotBasic()
	// fmt.Println("\n ---- tickers ----")
	// con2.TickBasic()
	// fmt.Println("\n ---- tickers ----")
	// con2.NewTicker()
	con2.ClosingChannel()
	con2.RangeChannel()
	fmt.Println("\n ---- Wait Group ----")
	con2.WaitGroups()

}

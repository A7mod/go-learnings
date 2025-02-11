package main

import (
	"fmt"

	"github.com/A7mod/go-learnings/basics"
	"github.com/A7mod/go-learnings/basics/arrays"
	"github.com/A7mod/go-learnings/basics/functions"
	"github.com/A7mod/go-learnings/basics/maps"
	"github.com/A7mod/go-learnings/basics/structs"
)

func main() {
	fmt.Println("Running Basics ... ")

	// basics i.e. main folder-ch
	fmt.Println("\n--- Basics ---")
	basics.ShowVariables()
	basics.ShowLoops()
	basics.CheckCondition()

	// arrays vale items
	fmt.Println("\n--- Arrays ---")
	arrays.ShowArrays()
	arrays.ShowSlices()
	arrays.ShowSlices2()

	// functions vale items
	fmt.Println("\n--- Functions ---")
	functions.ShowFunctions()
	functions.Represent()
	functions.NamedReturns()
	functions.ShowAnonfunc()
	functions.ShowVariadicFunc()
	functions.ShowFuncAsArgument()
	functions.ShowClosures()
	functions.ShowMethods()

	// maps vaale items
	fmt.Println("\n--- Maps ---")
	maps.ShowMaps()
	maps.ModifyMaps()
	maps.CheckKeyExists()
	maps.IterateMaps()
	maps.ShowNests()

	// Structs ka maamla
	fmt.Println("\n--- Structs ---")
	structs.StructBasics()
	structs.ShowRectangle()
	structs.ShowCar()
	structs.ShowUser()

}

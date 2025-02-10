package main

import (
	"fmt"

	"github.com/A7mod/go-learnings/basics"
	"github.com/A7mod/go-learnings/basics/arrays"
	"github.com/A7mod/go-learnings/basics/functions"
	"github.com/A7mod/go-learnings/basics/maps"
)

func main() {
	fmt.Println("Running Basics ... ")

	// basics i.e. main folder-ch
	basics.ShowVariables()
	basics.ShowLoops()
	basics.CheckCondition()

	// arrays vale items
	arrays.ShowArrays()
	arrays.ShowSlices()
	arrays.ShowSlices2()

	// functions vale items
	functions.Represent()
	functions.NamedReturns()
	functions.ShowAnonfunc()
	functions.ShowVariadicFunc()
	functions.ShowFunctions()
	functions.ShowFuncAsArgument()
	functions.ShowClosures()
	functions.ShowMethods()

	// maps vaale items
	maps.ShowMaps()
	maps.ModifyMaps()
	maps.CheckKeyExists()
	maps.IterateMaps()
	maps.ShowNests()

}

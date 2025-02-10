package main

import (
	"fmt"

	"github.com/A7mod/go-learnings/basics"
	"github.com/A7mod/go-learnings/basics/maps"
)

func main() {
	fmt.Println("Running Basics ... ")
	basics.ShowVariables()
	basics.ShowLoops()
	basics.ShowFunctions()
	basics.CheckCondition()
	basics.ShowArrays()
	basics.ShowSlices()
	basics.ShowSlices2()
	basics.Represent()
	basics.NamedReturns()
	basics.ShowAnonfunc()
	basics.ShowVariadicFunc()
	basics.ShowFuncAsArgument()
	basics.ShowClosures()
	basics.ShowMethods()

	//maps vaale items
	maps.ShowMaps()
	maps.ModifyMaps()
	maps.CheckKeyExists()

}

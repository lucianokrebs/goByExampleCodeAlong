package main

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goArrays"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goClosures"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goConstants"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goErrors"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goFor"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goFunctions"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goIfElse"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goInterfaces"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goMaps"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goMethods"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goMultipleReturns"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPointers"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goRange"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goRecursion"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goSlices"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goStructs"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goSwitch"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goValues"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goVariables"
)

func main() {
	fmt.Println("GoByExample code-along")
	goPrint.PrintDashes()

	goValues.Values()
	goVariables.Variables()
	goConstants.Constants()
	goFor.ForLoop()
	goIfElse.IfElse()
	goSwitch.GoSwitch()
	goArrays.GoArrays()
	goSlices.GoSlices()
	goMaps.GoMaps()
	goRange.GoRange()
	goFunctions.GoFunctions()
	goMultipleReturns.MultipleReturns()
	goClosures.ClosureExample()
	goRecursion.RecursionExample()
	goPointers.PointersExample()
	goStructs.StructExamples()
	goMethods.MethodExample()
	goInterfaces.InterfaceExamples()
	goErrors.ErrorsExample()
}

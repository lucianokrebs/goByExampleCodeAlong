package goFunctions

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func GoFunctions() {
	res := plus(2, 3)
	fmt.Println("res: ", res)

	resTwo := plusPlus(1, 2, 3)
	fmt.Println("resTwo: ", resTwo)

	goPrint.PrintDashes()
}

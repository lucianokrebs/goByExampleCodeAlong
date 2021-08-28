package goValues

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func Values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(!true)
	goPrint.PrintDashes()
}

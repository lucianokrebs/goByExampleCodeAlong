package goMultipleReturns

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func values() (int, int) {
	return 2, 3
}

func MultipleReturns() {
	a, b := values()
	fmt.Printf("a: %v | b %v\n", a, b)

	_, c := values()
	fmt.Printf("c: %v\n", c)

	goPrint.PrintDashes()
}

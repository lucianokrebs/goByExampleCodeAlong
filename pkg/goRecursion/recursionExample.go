package goRecursion

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func RecursionExample() {
	fmt.Println(fact(7))

	goPrint.PrintDashes()
}

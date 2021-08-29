package goClosures

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func ClosureExample() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newIntTwo := intSeq()
	fmt.Println(newIntTwo())

	goPrint.PrintDashes()
}

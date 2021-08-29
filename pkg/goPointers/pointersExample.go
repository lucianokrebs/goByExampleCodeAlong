package goPointers

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func zeroValue(val int) {
	val = 0
}

func zeroPointer(pointer *int) {
	*pointer = 0
}

func PointersExample() {
	i := 1
	fmt.Println("Initial: ", i)

	zeroValue(i)
	fmt.Println("zeroValue: ", i)

	zeroPointer(&i)
	fmt.Println("zeroPointer: ", i)

	goPrint.PrintDashes()
}

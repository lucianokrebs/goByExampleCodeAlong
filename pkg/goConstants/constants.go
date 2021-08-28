package goConstants

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func Constants() {
	const s string = "my constant"
	fmt.Printf("s: %v\n", s)

	const n int32 = 500000
	fmt.Printf("n: %v\n", n)

	goPrint.PrintDashes()
}

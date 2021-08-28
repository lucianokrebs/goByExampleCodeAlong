package goFor

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func ForLoop() {
	i := 1
	for i <= 3 {
		fmt.Printf("i: %v\n", i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Printf("j: %v\n", j)
	}

	for {
		fmt.Println("simple loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("n: %v\n", n)
	}

	goPrint.PrintDashes()
}

package goArrays

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func GoArrays() {
	// In Go, an array is a numbered sequence of elements of a specific length
	var a [5]int
	fmt.Printf("a: %v\n", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len(a): ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)

	goPrint.PrintDashes()
}

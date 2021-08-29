package goMaps

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func GoMaps() {
	m := make(map[string]int)
	fmt.Println("m: ", m)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("m: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("len: ", len(m))

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n: ", n)

	goPrint.PrintDashes()
}

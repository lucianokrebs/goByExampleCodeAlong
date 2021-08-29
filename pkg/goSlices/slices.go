package goSlices

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func GoSlices() {
	s := make([]string, 3)
	fmt.Println("s: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	fmt.Println("c before: ", c)
	copy(c, s)
	fmt.Println("c after: ", c)

	l := s[2:5]
	fmt.Println("l1: ", l)

	l = s[2:]
	fmt.Println("l2: ", l)

	l = s[:5]
	fmt.Println("l3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("t: ", t)

	goPrint.PrintDashes()
}

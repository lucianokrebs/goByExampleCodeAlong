package goRange

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

func GoRange() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		fmt.Printf("index: %v | num: %v\n", i, num)
		sum += num
	}
	fmt.Println("total sum: ", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("k: %v | v: %v\n", k, v)
	}
	for k := range kvs {
		fmt.Printf("k: %v \n", k)
	}
	for i, c := range "go" {
		fmt.Printf("i: %v | c: %v\n", i, c)
	}

	goPrint.PrintDashes()
}

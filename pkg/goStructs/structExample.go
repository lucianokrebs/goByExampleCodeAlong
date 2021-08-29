package goStructs

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func StructExamples() {
	fmt.Println(person{
		name: "Bob",
		age:  20,
	})
	fmt.Println(&person{
		name: "Ann",
		age:  30,
	})
	s := person{
		name: "Sean",
		age:  50,
	}

	j := newPerson("John")
	fmt.Println("j name: ", j.name)
	fmt.Println("j age: ", j.age)

	sp := &s
	fmt.Println("s: ", s.name)
	fmt.Println("sp: ", sp.age)
	sp.age = 51
	fmt.Println("sp: ", sp.age)

	goPrint.PrintDashes()
}

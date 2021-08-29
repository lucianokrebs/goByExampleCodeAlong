package goInterfaces

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func InterfaceExamples() {
	r := rect{
		width:  3,
		height: 4,
	}
	c := circle{
		radius: 5,
	}
	measure(r)
	measure(c)

	goPrint.PrintDashes()
}

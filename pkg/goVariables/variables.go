package goVariables

import "fmt"

func Variables() {
	var a = "initial"
	fmt.Printf("a: %v\n", a)

	var b, c = 1, 2
	fmt.Printf("b: %v, c: %v\n", b, c)

	var d = true
	fmt.Printf("d: %v\n", d)

	var e int
	fmt.Printf("e: %v\n", e)

	f := "apple"
	fmt.Printf("f: %v\n", f)

	fmt.Println("------------")
}

package main

import "fmt"

func constants() {
	const s string = "my constant"
	fmt.Printf("s: %v\n", s)

	const n int32 = 500000
	fmt.Printf("n: %v\n", n)

	fmt.Println("------------")
}

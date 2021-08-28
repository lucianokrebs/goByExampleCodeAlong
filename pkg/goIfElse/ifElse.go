package goIfElse

import (
	"crypto/rand"
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
	"log"
	"math/big"
)

func IfElse() {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatal("error creating random number: ", err)
	}

	r := randomNumber.Int64()
	fmt.Printf("r: %v\n", r)

	if r < int64(33) {
		fmt.Println("smaller than 33")
	} else if r >= int64(33) && r <= int64(66) {
		fmt.Println("between 33 and 66")
	} else {
		fmt.Println("bigger than 66")
	}

	goPrint.PrintDashes()
}

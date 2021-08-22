package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func ifElse() {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatal("error creating random number: ", err)
	}

	r := randomNumber.Int64()
	fmt.Printf("r: %v\n", r)

	if r < int64(33) {
		fmt.Printf("smaller than 33")
	} else if r >= int64(33) && r <= int64(66) {
		fmt.Printf("between 33 and 66")
	} else {
		fmt.Printf("bigger than 66")
	}
}

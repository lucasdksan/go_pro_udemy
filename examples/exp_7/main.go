package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := bcrypt.GenerateFromPassword([]byte("123456789"), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(hash))

	if err = bcrypt.CompareHashAndPassword(hash, []byte("1234567890")); err != nil {
		fmt.Println("if", err)
	} else {
		fmt.Println("else", err)
	}
}

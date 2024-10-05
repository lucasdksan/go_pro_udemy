package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	fmt.Println(port)
}

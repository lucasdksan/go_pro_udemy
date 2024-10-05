package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "9000", "Server Port")

	flag.Parse()

	fmt.Println("Server is running in port", port)
}

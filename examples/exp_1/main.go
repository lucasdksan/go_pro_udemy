package main

import (
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("index.html")

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, "Lucas Silva")

	if err != nil {
		panic(err)
	}
}

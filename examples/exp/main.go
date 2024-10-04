package main

import (
	"os"
	"text/template"
)

func main() {
	t, err := template.New("teste").Parse("<h1>Este é um título, {{ . }}</h1>")

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, "Lucas da Silva")

	if err != nil {
		panic(err)
	}
}

package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

func render(w http.ResponseWriter, page string, data interface{}, status int) error {
	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
	}

	files = append(files, fmt.Sprintf("views/templates/%s", page))

	t, err := template.ParseFiles(files...)

	if err != nil {
		return ErrorInternalServer("error executing this page")
	}

	buff := &bytes.Buffer{}

	if err = t.ExecuteTemplate(buff, "layout", data); err != nil {
		return ErrorInternalServer("error in template")
	}

	w.WriteHeader(status)
	buff.WriteTo(w)

	return nil
}

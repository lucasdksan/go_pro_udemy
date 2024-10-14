package handlers

import (
	"bytes"
	"net/http"
	"text/template"
)

func render(w http.ResponseWriter, page string, data interface{}) error {
	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
	}

	files = append(files, "views/templates/"+page)

	t, err := template.ParseFiles(files...)

	if err != nil {
		return ErrorInternalServer("error executing this page")
	}

	buff := &bytes.Buffer{}

	if err = t.ExecuteTemplate(buff, "layout", data); err != nil {
		return ErrorInternalServer("error in template")
	}

	buff.WriteTo(w)

	return nil
}

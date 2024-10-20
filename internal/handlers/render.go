package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func render(w http.ResponseWriter, r *http.Request, page string, data interface{}, status int) error {
	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
	}

	files = append(files, fmt.Sprintf("views/templates/%s", page))
	t := template.New("").Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
		"csrfToken": func() string {
			return csrf.Token(r)
		},
	})

	t, err := t.ParseFiles(files...)

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

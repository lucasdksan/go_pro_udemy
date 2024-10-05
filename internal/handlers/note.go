package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type noteHandler struct{}

func NewNoteHandler() *noteHandler {
	return &noteHandler{}
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
		"views/templates/home.html",
	}
	t, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(w, "Error executing this page", http.StatusInternalServerError)
	}

	t.ExecuteTemplate(w, "layout", nil)
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
		"views/templates/note-view.html",
	}

	if id == "" {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(w, "Error executing this page", http.StatusInternalServerError)
	}

	t.ExecuteTemplate(w, "layout", id)
}

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criar uma Nota")
}

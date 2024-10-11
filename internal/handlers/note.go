package handlers

import (
	"fmt"
	"go_pro/internal/repositories"
	"net/http"
	"strconv"
	"text/template"
)

type noteHandler struct {
	repo repositories.NoteRepository
}

func NewNoteHandler(repo repositories.NoteRepository) *noteHandler {
	return &noteHandler{repo: repo}
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return ErrorNotFound("page not found")
	}

	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
		"views/templates/home.html",
	}
	t, err := template.ParseFiles(files...)

	if err != nil {
		return ErrorInternalServer("Error executing this page")
	}

	if err = t.ExecuteTemplate(w, "layout", nil); err != nil {
		return ErrorInternalServer("error in template")
	}

	return nil
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("id")
	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
		"views/templates/note-view.html",
	}

	if idParam == "" {
		return ErrorBadRequest("note not found")
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err
	}

	t, err := template.ParseFiles(files...)

	if err != nil {
		return ErrorInternalServer("error executing this page")
	}

	note, err := nh.repo.GetById(id)

	if err != nil {
		return err
	}

	if err = t.ExecuteTemplate(w, "layout", note); err != nil {
		return ErrorInternalServer("error in template")
	}

	return nil
}

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criar uma Nota")
}

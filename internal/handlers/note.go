package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"go_pro/internal/apperrors"
	"go_pro/internal/dtos"
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

	notes, err := nh.repo.List(r.Context())

	if err != nil {
		return err
	}

	if err = t.ExecuteTemplate(w, "layout", dtos.NewNoteResponseFromNoteList(notes)); err != nil {
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

	note, err := nh.repo.GetById(r.Context(), id)

	if err != nil {
		return err
	}

	buff := &bytes.Buffer{}

	if err = t.ExecuteTemplate(buff, "layout", dtos.NewNoteResponseFromNote(note)); err != nil {
		return ErrorInternalServer("error in template")
	}

	buff.WriteTo(w)

	return nil
}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) error {
	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
		"views/templates/note-new.html",
	}

	t, err := template.ParseFiles(files...)

	if err != nil {
		return ErrorInternalServer("error executing this page")
	}

	if err = t.ExecuteTemplate(w, "layout", dtos.NewNoteRequest()); err != nil {
		return ErrorInternalServer("error in template")
	}

	return nil
}

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		return apperrors.NewWithStatus(errors.New("operation not permitted"), http.StatusInternalServerError)
	}

	if err := r.ParseForm(); err != nil {
		return err
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	color := r.PostForm.Get("color")

	note, err := nh.repo.Create(r.Context(), title, content, color)

	if err != nil {
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/notes/view?id=%d", note.Id.Int), http.StatusSeeOther)

	return nil
}

package handlers

import (
	"errors"
	"fmt"
	"go_pro/internal/apperrors"
	"go_pro/internal/dtos"
	"go_pro/internal/repositories"
	"net/http"
	"strconv"
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

	notes, err := nh.repo.List(r.Context())

	if err != nil {
		return err
	}

	if err = render(w, "home.html", dtos.NewNoteResponseFromNoteList(notes)); err != nil {
		return ErrorInternalServer("error in template")
	}

	return nil
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		return ErrorBadRequest("note not found")
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err
	}

	note, err := nh.repo.GetById(r.Context(), id)

	if err != nil {
		return err
	}

	if err = render(w, "note-view.html", dtos.NewNoteCombined(note)); err != nil {
		return ErrorInternalServer("error in template")
	}

	return nil
}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) error {
	if err := render(w, "note-new.html", dtos.NewNoteRequest()); err != nil {
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

func (nh *noteHandler) NoteDelete(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodPost)

		return apperrors.NewWithStatus(errors.New("operation not permitted"), http.StatusInternalServerError)
	}

	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		return ErrorBadRequest("note not found")
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err
	}

	err = nh.repo.Delete(r.Context(), id)

	if err != nil {
		return ErrorInternalServer("Error deleting note")
	}

	return nil
}

func (nh *noteHandler) NoteEdit(w http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		return ErrorBadRequest("note not found")
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	color := r.PostForm.Get("color")

	note, err := nh.repo.Update(r.Context(), id, title, content, color)

	if err != nil {
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/notes/view?id=%d", note.Id.Int), http.StatusSeeOther)

	return nil
}

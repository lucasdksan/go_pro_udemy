package handlers

import (
	"fmt"
	"go_pro/internal/apperrors"
	"go_pro/internal/dtos"
	"go_pro/internal/models"
	"go_pro/internal/render"
	"go_pro/internal/repositories"
	"net/http"
	"strconv"
	"strings"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/csrf"
)

type noteHandler struct {
	render  *render.RenderTemplate
	session *scs.SessionManager
	repo    repositories.NoteRepository
}

func NewNoteHandler(render *render.RenderTemplate, session *scs.SessionManager, repo repositories.NoteRepository) *noteHandler {
	return &noteHandler{repo: repo, render: render, session: session}
}

func (nh *noteHandler) getUserIdFromSession(r *http.Request) int64 {
	return nh.session.GetInt64(r.Context(), "userId")
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) error {
	notes, err := nh.repo.List(r.Context(), int(nh.getUserIdFromSession(r)))

	if err != nil {
		return err
	}

	if err = nh.render.RenderPage(w, r, "note-home.html", dtos.NewNoteResponseFromNoteList(notes), http.StatusOK); err != nil {
		return err
	}

	return nil
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	idParam := r.PathValue("id")

	if idParam == "" {
		return apperrors.ErrorBadRequest("note not found")
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err
	}

	note, err := nh.repo.GetById(r.Context(), int(nh.getUserIdFromSession(r)), id)

	if err != nil {
		return err
	}

	if err = nh.render.RenderPage(w, r, "note-view.html", dtos.NewNoteResponseFromNote(note), http.StatusOK); err != nil {
		return err
	}

	return nil
}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) error {
	data := dtos.NewNoteRequest(nil)

	data.CSRFField = csrf.TemplateField(r)

	if err := nh.render.RenderPage(w, r, "note-new.html", data, http.StatusOK); err != nil {
		return err
	}

	return nil
}

func (nh *noteHandler) NoteSave(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	idParam := r.PostForm.Get("id")
	id, _ := strconv.Atoi(idParam)
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	color := r.PostForm.Get("color")
	data := dtos.NewNoteRequest(nil)

	data.Color = color
	data.Content = content
	data.Title = title
	data.Id = id

	if strings.TrimSpace(content) == "" {
		data.AddFieldError("content", "content is required")
	}

	if !data.Valid() {
		if id > 0 {
			nh.render.RenderPage(w, r, "note-edit.html", data, http.StatusUnprocessableEntity)
		} else {
			nh.render.RenderPage(w, r, "note-new.html", data, http.StatusUnprocessableEntity)
		}

		return nil
	}

	var note *models.Note
	var err error

	if id > 0 {
		note, err = nh.repo.Update(r.Context(), int(nh.getUserIdFromSession(r)), id, title, content, color)
	} else {
		note, err = nh.repo.Create(r.Context(), int(nh.getUserIdFromSession(r)), title, content, color)
	}

	if err != nil {
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/notes/%d", note.Id.Int), http.StatusSeeOther)

	return nil
}

func (nh *noteHandler) NoteDelete(w http.ResponseWriter, r *http.Request) error {
	idParam := r.PathValue("id")

	if idParam == "" {
		return apperrors.ErrorBadRequest("note not found")
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err
	}

	err = nh.repo.Delete(r.Context(), int(nh.getUserIdFromSession(r)), id)

	if err != nil {
		return apperrors.ErrorInternalServer("Error deleting note")
	}

	return nil
}

func (nh *noteHandler) NoteEdit(w http.ResponseWriter, r *http.Request) error {
	idParam := r.PathValue("id")

	if idParam == "" {
		return apperrors.ErrorBadRequest("note not found")
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err
	}

	note, err := nh.repo.GetById(r.Context(), int(nh.getUserIdFromSession(r)), id)

	if err != nil {
		return err
	}

	if err = nh.render.RenderPage(w, r, "note-edit.html", dtos.NewNoteRequest(note), http.StatusOK); err != nil {
		return err
	}

	return nil
}

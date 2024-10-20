package dtos

import (
	"fmt"
	"go_pro/internal/models"
	"go_pro/internal/validations"
	"html/template"
)

type NoteResponse struct {
	Id      int
	Title   string
	Content string
	Color   string
}

type NoteRequest struct {
	Id        int
	Title     string
	Content   string
	Color     string
	Colors    []string
	CSRFField template.HTML
	validations.FormValidator
}

func NewNoteRequest(note *models.Note) (req NoteRequest) {
	for i := 1; i <= 9; i++ {
		req.Colors = append(req.Colors, fmt.Sprintf("color%d", i))
	}

	if note != nil {
		req.Id = int(note.Id.Int.Int64())
		req.Title = note.Title.String
		req.Color = note.Color.String
		req.Content = note.Content.String
	} else {
		req.Color = "color1"
	}

	return
}

func NewNoteResponseFromNote(note *models.Note) (res NoteResponse) {
	res.Id = int(note.Id.Int.Int64())
	res.Title = note.Title.String
	res.Content = note.Content.String
	res.Color = note.Color.String

	return
}

func NewNoteResponseFromNoteList(notes []models.Note) (res []NoteResponse) {
	for _, note := range notes {
		res = append(res, NewNoteResponseFromNote(&note))
	}

	return
}

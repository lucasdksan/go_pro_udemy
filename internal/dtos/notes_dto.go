package dtos

import (
	"fmt"
	"go_pro/internal/models"
)

type NoteResponse struct {
	Id      int
	Title   string
	Content string
	Color   string
}

type NoteRequest struct {
	Id      int
	Title   string
	Content string
	Color   string
	Colors  []string
}

func NewNoteRequest() (req NoteRequest) {
	req.Color = "color1"

	for i := 1; i <= 9; i++ {
		req.Colors = append(req.Colors, fmt.Sprintf("color%d", i))
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

func NewNoteCombined(note *models.Note) (req NoteRequest) {
	req.Id = int(note.Id.Int.Int64())
	req.Title = note.Title.String
	req.Content = note.Content.String
	req.Color = note.Color.String

	for i := 1; i <= 9; i++ {
		req.Colors = append(req.Colors, fmt.Sprintf("color%d", i))
	}

	return
}

func NewNoteResponseFromNoteList(notes []models.Note) (res []NoteResponse) {
	for _, note := range notes {
		res = append(res, NewNoteResponseFromNote(&note))
	}

	return
}

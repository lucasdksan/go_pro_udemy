package dtos

import "go_pro/internal/models"

type NoteResponse struct {
	Id      int
	Title   string
	Content string
	Color   string
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

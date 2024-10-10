package database

var ListNoteQuery string = `
	select id, title, content, color, created_at from notes;
`

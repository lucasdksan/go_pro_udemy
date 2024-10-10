package querys

var (
	ListNoteQuery string = `
		select id, title, content, color, created_at from notes;
	`
	GetByIdNoteQuery string = `
		select * from notes where id=$1
	`
	CreateNoteQuery string = `
		insert into notes (title, content, color)
		values ($1, $2, $3)
		returning id, created_at
	`
	UpdateNoteQuery string = `
		update notes set title = $1, content = coalesce($2, content), color = $3, updated_at = $4 where id = $5;
	`
	DeleteNoteQuery string = `
		delete from notes where id = $1
	`
)

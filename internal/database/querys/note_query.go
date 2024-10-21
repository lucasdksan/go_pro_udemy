package querys

var (
	ListNoteQuery string = `
		select id, title, content, color, created_at, updated_at from notes where user_id = $1;
	`
	GetByIdNoteQuery string = `
		select id, title, content, color, created_at, updated_at from notes where id = $1 and user_id = $2;
	`
	CreateNoteQuery string = `
		INSERT INTO notes (title, content, color, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at;
	`
	UpdateNoteQuery string = `
		update notes set title = $1, content = coalesce($2, content), color = $3, updated_at = $4 where id = $5 and user_id = $6;
	`
	DeleteNoteQuery string = `
		delete from notes where id = $1 and user_id = $2;
	`
)

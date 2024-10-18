package querys

var (
	CreateTokenQuery string = `
		insert into users_confirmation_tokens (user_id, token)
		values ($1, $2) 
		returning id, created_at;
	`
	UpdateTokenConfirmedQuery string = `
		update users_confirmation_tokens set confirmed = true, updated_at = now() where id = $1;
	`
)

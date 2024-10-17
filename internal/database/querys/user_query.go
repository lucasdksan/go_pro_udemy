package querys

var (
	CreateUserQuery string = `
		insert into users (email, password)
		values ($1, $2)
		returning id, created_at
	`
)

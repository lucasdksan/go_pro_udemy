package querys

var (
	CreateUserQuery string = `
		insert into users (email, password)
		values ($1, $2)
		returning id, created_at;
	`
	GetUserJoinUserTokenQuery string = `
		select u.id u_id, t.id t_id from users u inner join users_confirmation_tokens t
		on u.id = t.user_id
		where u.active = false
		and t.confirmed = false
		and t.token = $1;
	`
	UpdateUserConfirmedQuery string = `
		update users set active = true, updated_at = now() where id = $1;
	`
	FindByEmailQuery string = `
		select id, email, password, active from users where email = $1;
	`
)

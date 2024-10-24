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
	SelectPasswordByTokenQuery string = `
		select u.id u_id, u.email, t.id t_id from users u inner join users_confirmation_tokens t
		on u.id = t.user_id
		where t.confirmed = false
		and t.token = $1;
	`
	UpdatePasswordByTokenQuery string = `
		update users_confirmation_tokens set confirmed = true, updated_at = now()
		where id = $1;
	`
	UpdatePasswordQuery string = `
		update users set password = $1, updated_at = now() where id = $2;
	`
)

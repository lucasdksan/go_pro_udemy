create table if not exists users_confirmation_tokens(
    id bigserial primary key,
    user_id bigserial not null references users(id) on delete cascade,
    token text not null,
    confirmed boolean not null default false,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
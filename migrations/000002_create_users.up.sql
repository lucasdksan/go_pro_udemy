create table if not exists users(
    id bigserial primary key,
    email text not null unique,
    password text not null,
    active boolean not null default false,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
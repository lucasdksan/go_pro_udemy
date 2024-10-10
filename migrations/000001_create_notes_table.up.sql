create table if not exists notes (
    id bigserial primary key,
    title text not null,
    content text,
    color varchar(100) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
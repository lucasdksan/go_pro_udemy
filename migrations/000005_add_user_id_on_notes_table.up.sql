alter table notes
add column user_id bigint not null references users(id)
on delete cascade;
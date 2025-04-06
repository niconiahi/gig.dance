-- +goose Up
-- +goose StatementBegin
create table user (
    id integer not null primary key autoincrement,
    name text not null,
    email text not null
);

create table session (
    id text not null primary key,
    expires_at integer not null,
    user_id integer not null,
    foreign key (user_id) references user (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
drop table session;
-- +goose StatementEnd

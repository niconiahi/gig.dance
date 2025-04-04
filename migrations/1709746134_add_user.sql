-- +goose Up
-- +goose StatementBegin
create table user (
    id integer primary key autoincrement,
    username text,
    name text,
    surname text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd

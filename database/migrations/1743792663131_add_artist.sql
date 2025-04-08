-- +goose Up
-- +goose StatementBegin
create table artist (
    id integer primary key autoincrement,
    name text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table artist;
-- +goose StatementEnd

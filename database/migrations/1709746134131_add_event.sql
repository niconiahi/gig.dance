-- +goose Up
-- +goose StatementBegin
create table event (
    id integer primary key autoincrement,
    name text not null,
    url text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table event;
-- +goose StatementEnd

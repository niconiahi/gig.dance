-- +goose Up
-- +goose StatementBegin
create table location (
    id integer primary key autoincrement,
    place_id text not null,
    name text not null,
    address text not null,
    latitude text not null,
    longitude text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table location;
-- +goose StatementEnd

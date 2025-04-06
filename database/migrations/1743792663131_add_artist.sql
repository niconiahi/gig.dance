-- +goose Up
-- +goose StatementBegin
create table artist (
    id integer primary key autoincrement,
    name text not null
);
create table event_artist (
    id integer primary key autoincrement,
    event_id integer not null,
    artist_id integer not null,
    foreign key (event_id) references event(id),
    foreign key (artist_id) references artist(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table artist;
drop table event_artist;
-- +goose StatementEnd

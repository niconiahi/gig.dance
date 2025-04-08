-- +goose Up
-- +goose StatementBegin
create table event (
    id integer primary key autoincrement,
    url text not null,
    "date" text not null,
    name text not null,
    location_id integer not null,
    foreign key (location_id) references location (id)
);
create table event_artist (
    id integer primary key autoincrement,
    event_id integer not null,
    artist_id integer not null,
    foreign key (event_id) references event (id),
    foreign key (artist_id) references artist (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table event;
drop table event_artist;
-- +goose StatementEnd

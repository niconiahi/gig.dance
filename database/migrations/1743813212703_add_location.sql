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
create table event_location (
    id integer primary key autoincrement,
    event_id integer not null,
    location_id integer not null,
    foreign key (event_id) references event(id),
    foreign key (location_id) references location(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table location;
drop table event_location;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
create table manager (
    id integer primary key autoincrement,
    user_id integer not null,
    foreign key (user_id) references user (id)
);
create table artist_manager (
    id integer primary key autoincrement,
    artist_id integer not null,
    manager_id integer not null,
    foreign key (artist_id) references artist (id),
    foreign key (manager_id) references manager (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table manager;
drop table artist_manager;
-- +goose StatementEnd

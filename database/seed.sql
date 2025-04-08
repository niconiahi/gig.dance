-- add manager's user
insert into user (name, email)
values ('Alice', 'manager@gig.dance');

with manager_user as (
    select id
    from user
    where email = 'manager@gig.dance'
)

-- add manager
insert into manager (user_id)
select id from manager_user;

-- add artist
insert into artist (name)
values ('Rafa NPM');

-- add artist manager
with rafa_manager as (
    select manager.id
    from manager
    inner join user on manager.user_id = user.id
    where user.email = 'manager@gig.dance'
),

rafa_artist as (
    select id
    from artist
    where name = 'Rafa NPM'
)

insert into artist_manager (artist_id, manager_id)
select
    rafa_artist.id as artist_id,
    rafa_manager.id as manager_id
from rafa_artist, rafa_manager;

-- add location
insert into location (place_id, name, address, longitude, latitude)
values (
    'ChIJkwq0b-u1vJURlli1Yhuxyh8',
    'Movistar Arena',
    'Humboldt 450, C1414CSJ Cdad. Autónoma de Buenos Aires, Argentina',
    '-34.5945093',
    '-58.4480012'
);

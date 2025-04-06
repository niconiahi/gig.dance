insert into user (name, email)
values ('Alice', 'manager@gig.dance');

insert into location (name, email)
values (
    'ChIJkwq0b-u1vJURlli1Yhuxyh8',
    'Movistar Arena',
    'Humboldt 450, C1414CSJ Cdad. Autónoma de Buenos Aires, Argentina',
    '-34.5945093',
    '-58.4480012'
);

with manager_user as (
    select id
    from user
    where email = 'manager@gig.dance'
)

insert into manager (user_id)
select id from manager_user;

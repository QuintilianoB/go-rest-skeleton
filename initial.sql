create table user
(
    id       int auto_increment,
    user     varchar(256) not null,
    password varchar(256) not null,
    constraint user_pk
        primary key (id)
);

create unique index user_user_uindex
    on user (user);


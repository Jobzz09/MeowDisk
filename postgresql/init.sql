CREATE USER meow_admin;
CREATE DATABASE meow_disk OWNER meow_admin;

CREATE TABLE meow_disk.user_data (
    id       int primary key,
    login    varchar(30) not null,
    password varchar(30) not null
);

CREATE TABLE mew_disk.file_data(
    id INT PRIMARY KEY,
    name varchar(150) not null,
    type varchar(150) not null,
    hashsum varchar(150) not null
);


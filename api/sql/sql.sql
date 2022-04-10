CREATE DATABASE IF NOT EXISTS sgamer;
USE sgamer;

DROP TABLE IF EXISTS user;

CREATE TABLE user(
    id int auto_increment primary key,
    username varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    pass varchar(200) not null,
    createDate timestamp default current_timestamp()
) ENGINE=INNODB;

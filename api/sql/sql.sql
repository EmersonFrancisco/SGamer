CREATE DATABASE IF NOT EXISTS sgamer;
USE sgamer;

DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS follower;

CREATE TABLE user(
    id int auto_increment primary key,
    username varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    pass varchar(200) not null,
    createDate timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE follower(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES user(id)
    on DELETE CASCADE,
    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES user(id)
    ON DELETE CASCADE,

    primary key(user_id,follower_id)
) ENGINE=INNODB;
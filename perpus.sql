CREATE DATABASE perpus;

USE perpus;

CREATE TABLE user(
id int PRIMARY KEY AUTO_INCREMENT,
nama varchar(100),
address varchar(50),
hp varchar(13),
password varchar(10)
);

CREATE TABLE genre(
id int PRIMARY KEY AUTO_INCREMENT,
nama varchar(20)
);

CREATE TABLE book(
id int PRIMARY KEY AUTO_INCREMENT,
user_id int,
genre_id int,
title varchar(100),
isbn varchar(50),
author varchar(100),
penerbit varchar(100),
jumlah int,
deskripsi varchar(100),
CONSTRAINT fk_userbook FOREIGN KEY (user_id) REFERENCES user(id),
CONSTRAINT fk_genrebook FOREIGN KEY (genre_id) REFERENCES genre(id)
);

CREATE TABLE rent(
id int PRIMARY KEY AUTO_INCREMENT,
user_id int,
book_id int,
created_at timestamp,
return_date date,
CONSTRAINT fk_user_rent FOREIGN KEY (user_id) REFERENCES user(id),
CONSTRAINT fk_book_rent FOREIGN KEY (book_id) REFERENCES book(id)
);

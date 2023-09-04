CREATE USER 'authuser'@'%' IDENTIFIED BY 'mypass';

GRANT ALL PRIVILEGES ON auth.* TO 'authuser'@'%';

FLUSH PRIVILEGES;

USE auth;

CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) UNIQUE NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO users(email, password) VALUES("ratest@gmail.com", "$2a$12$UwYb.OqXxFdTC/ZOxxKyWO4XHlf9yMBircaIafBmuLAeFHwVObVwu");
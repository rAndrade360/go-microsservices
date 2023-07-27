CREATE USER 'authuser'@'localhost' IDENTIFIED BY 'mypass';

GRANT ALL PRIVILEGES ON auth.* TO 'authuser'@'localhost';

FLUSH PRIVILEGES;

USE auth;

CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) UNIQUE NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO users(email, password) VALUES("ratest@gmail.com", "$2a$10$tHuykfXzTgTZgLiVY57Ep.bA57dqbZPnc5XnNWsFHYINTvZzvpuhy");
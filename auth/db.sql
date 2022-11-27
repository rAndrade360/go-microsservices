CREATE USER 'authuser'@'172.17.0.1' IDENTIFIED BY 'mypass';

GRANT ALL PRIVILEGES ON auth.* TO 'authuser'@'172.17.0.1';

FLUSH PRIVILEGES;

USE auth;

CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) UNIQUE NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO users(email, password) VALUES("renandotcorrea@gmail.com", "$2a$10$LQliu1B8/S2xH1fdWW7ZC.oUefxsrRbhKEunPPa/xzY0E3y13IZt.");
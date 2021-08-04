CREATE DATABASE todos CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE todos;

CREATE TABLE user (
                       id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       hashed_password CHAR(60) NOT NULL,
                       created DATETIME NOT NULL,
                       active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE user ADD CONSTRAINT user_uc_email UNIQUE (email);

CREATE TABLE todo (
                          id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
                          title VARCHAR(100) NOT NULL,
                          content TEXT NOT NULL,
                          created DATETIME NOT NULL,
                          completion_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          completed BOOLEAN NOT NULL,
                          user_id INTEGER,
                          FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE INDEX idx_todo_created ON todo(created);
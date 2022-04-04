CREATE TABLE user (
                       id INTEGER NOT NULL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       hashed_password CHAR(60) NOT NULL,
                       created DATETIME NOT NULL,
                       active BOOLEAN NOT NULL DEFAULT TRUE,
                       UNIQUE(email)
);

CREATE TABLE todo (
                          id INTEGER NOT NULL PRIMARY KEY,
                          title VARCHAR(100) NOT NULL,
                          content TEXT NOT NULL,
                          created DATETIME NOT NULL,
                          completion_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          completed BOOLEAN NOT NULL,
                          user_id INTEGER,
                          FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE INDEX idx_todo_created ON todo(created);

CREATE TABLE tasks (
    id VARCHAR(255) UNIQUE,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(500)
)
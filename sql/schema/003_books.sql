-- +goose Up
CREATE TABLE books (
    id UUID PRIMARY KEY,
    isbn VARCHAR(13) NOT NULL UNIQUE,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    num_pages INT,
    price INT
);



-- +goose Down
DROP TABLE books;
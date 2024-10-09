-- +goose Up
CREATE TABLE books (
    id UUID PRIMARY KEY,
    isbn VARCHAR(13) NOT NULL UNIQUE,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    num_pages INT NOT NULL,
    price INT NOT NULL
);



-- +goose Down
DROP TABLE books;
-- +goose Up
CREATE TABLE authors (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
);

--  +goose Down
DROP TABLE authors;
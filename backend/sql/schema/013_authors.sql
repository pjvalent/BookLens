-- +goose Up
CREATE TABLE authors (
    id UUID PRIMARY KEY,
    author_name TEXT NOT NULL
);

-- +goose Down
DROP TABLE authors;
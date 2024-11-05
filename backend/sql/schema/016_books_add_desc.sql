-- +goose Up
ALTER TABLE books ADD COLUMN IF NOT EXISTS book_desc TEXT;

-- +goose Down
ALTER TABLE books DROP COLUMN book_desc;
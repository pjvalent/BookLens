-- +goose Up
ALTER TABLE books ADD COLUMN author_id BIGINT;

-- +goose Down
ALTER TABLE books DROP COLUMN author_id;
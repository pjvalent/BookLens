-- +goose Up
ALTER TABLE books ADD COLUMN genres jsonb;

-- +goose Down
ALTER TABLE book DROP COLUMN genres;
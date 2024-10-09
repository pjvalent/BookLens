-- +goose Up
ALTER TABLE books ADD COLUMN created_at TIMESTAMP NOT NULL;
ALTER TABLE books ADD COLUMN updated_at TIMESTAMP NOT NULL;


-- +goose Down
ALTER TABLE books DROP COLUMN created_at;
ALTER TABLE books DROP COLUMN updated_at;


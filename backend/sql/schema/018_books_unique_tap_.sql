-- +goose Up
ALTER TABLE books ADD CONSTRAINT unique_title_publisher UNIQUE (title, publisher);

-- +goose Down
ALTER TABLE books DROP CONSTRAINT unique_title_publisher;
-- +goose Up
ALTER TABLE books_generes ADD COLUMN new_id UUID;

ALTER TABLE books_generes DROP COLUMN id;

ALTER TABLE books_generes RENAME COLUMN new_id TO id;

-- +goose Down
ALTER TABLE books_generes ALTER COLUMN id TYPE SERIAL;
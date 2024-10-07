-- +goose Up
ALTER TABLE books ADD CONSTRAINT title_author_unique UNIQUE(title, author);


-- +goose Down
ALTER TABLE books DROP CONSTRAINT title_author_unique;

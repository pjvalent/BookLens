-- +goose Up
ALTER TABLE books ADD CONSTRAINT fk_books_author FOREIGN KEY (author_id) REFERENCES authors(id);

-- +goose Down
ALTER TABLE books DROP CONSTRAINT fk_books_author;
-- +goose Up
ALTER TABLE books ADD CONSTRAINT fk_books_author FOREIGN KEY (author_id) REFERENCES authors(id);

-- +goose Down
ALTER TABLE bokos DROP CONSTRAINT fk_books_author;
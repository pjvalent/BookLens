-- +goose Up
CREATE TABLE authors (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
);


ALTER TABLE books 
ADD COLUMN publication_day SMALLINT,
ADD COLUMN publication_month SMALLINT,
ADD COLUMN pubilcation_year SMALLINT,
ADD COLUMN publisher TEXT,
ADD COLUMN book_desc TEXT,
ADD COLUMN format TEXT,
ADD COLUMN author_id UUID,
ADD CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES authors(id);


--  +goose Down
DROP TABLE authors;
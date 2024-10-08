-- +goose Up

CREATE TABLE books_generes (
    id SERIAL PRIMARY KEY,
    isbn VARCHAR(13) NOT NULL,
    genere_id UUID NOT NULL,
    UNIQUE(isbn, genere_id),
    FOREIGN KEY (isbn) REFERENCES books(isbn) ON DELETE CASCADE,
    FOREIGN KEY (genere_id) REFERENCES generes(genere_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE books_generes;
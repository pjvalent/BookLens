-- +goose Up
CREATE TABLE book_embeddings (
    id UUID PRIMARY KEY,
    book_id UUID NOT NULL,
    embedding vector(384),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE book_embeddings;

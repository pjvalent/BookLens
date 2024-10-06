-- +goose Up
CREATE TABLE reviews (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    book_id UUID NOT NULL,
    rating INT NOT NULL,
    review_text TEXT,
    spoiler_tag BOOLEAN,
    UNIQUE(user_id, book_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE reviews;
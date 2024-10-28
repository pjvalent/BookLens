-- name: CreateBook :one
INSERT INTO books (id, isbn, created_at, updated_at, title, author, num_pages, price)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetBookByTitleAuthor :one
SELECT * FROM books WHERE title=$1 AND author=$2;

-- name: SimilarBooksByDesc :many
SELECT b.title, b.publisher, b.book_desc
FROM books b
INNER JOIN book_embeddings be ON b.id = be.book_id
CROSS JOIN (
    SELECT embedding
    FROM book_embeddings
    WHERE book_id = (SELECT id FROM books WHERE books.isbn = $1)
) AS target_embedding
WHERE b.id != (SELECT id FROM books WHERE isbn = $1)
ORDER BY be.embedding <=> target_embedding.embedding
LIMIT $2;

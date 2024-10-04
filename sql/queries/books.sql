-- name: CreateBook :one
INSERT INTO books (id, isbn, created_at, updated_at, title, author, num_pages, price)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

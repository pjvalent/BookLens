-- name: CreateBook :one
INSERT INTO books (isbn, title, author, num_pages, price)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

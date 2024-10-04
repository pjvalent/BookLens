-- name: CreateBookGenere :one
INSERT INTO books_generes (isbn, genere_id)
VALUES ($1, $2)
RETURNING *;
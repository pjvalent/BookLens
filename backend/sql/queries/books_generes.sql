-- name: CreateBooksGeneres :one
INSERT INTO books_generes (id, isbn, genere_id)
VALUES $1, $2, $3
RETURNING *;
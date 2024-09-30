-- name: CreateGenere :one
INSERT INTO generes (name)
VALUES ($1)
RETURNING *;
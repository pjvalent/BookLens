-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, first_name, last_name, email)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
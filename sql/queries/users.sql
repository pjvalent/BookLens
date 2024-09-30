-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, first_name, last_name, email, account_balance, api_key)
VALUES ($1, $2, $3, $4, $5, $6, $7, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;


-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;
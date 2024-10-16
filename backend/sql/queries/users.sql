-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, first_name, last_name, email, account_balance, api_key, user_password)
VALUES ($1, $2, $3, $4, $5, $6, $7, encode(sha256(random()::text::bytea), 'hex'), $8)
RETURNING *;


-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: DeleteUserByUserID :exec
DELETE FROM users WHERE id=$1;
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, first_name, last_name, email, account_balance, api_key)
VALUES ($1, $2, $3, $4, $5, $6, $7, encode(sha256(random()::text::bytea), 'hex'))
RETURNING id, created_at, updated_at, first_name, last_name, email, account_balance, api_key
`

type CreateUserParams struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	FirstName      string
	LastName       string
	Email          string
	AccountBalance int64
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.AccountBalance,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.AccountBalance,
		&i.ApiKey,
	)
	return i, err
}

const deleteUserByUserID = `-- name: DeleteUserByUserID :exec
DELETE FROM users WHERE id=$1
`

func (q *Queries) DeleteUserByUserID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserByUserID, id)
	return err
}

const getUserByApiKey = `-- name: GetUserByApiKey :one
SELECT id, created_at, updated_at, first_name, last_name, email, account_balance, api_key FROM users WHERE api_key = $1
`

func (q *Queries) GetUserByApiKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByApiKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.AccountBalance,
		&i.ApiKey,
	)
	return i, err
}
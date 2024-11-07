// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: authors.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (id, author_name, average_rating, author_id, text_review_count, ratings_count)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, author_name, average_rating, author_id, text_review_count, ratings_count
`

type CreateAuthorParams struct {
	ID              uuid.UUID
	AuthorName      string
	AverageRating   sql.NullString
	AuthorID        sql.NullInt64
	TextReviewCount sql.NullInt64
	RatingsCount    sql.NullInt64
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor,
		arg.ID,
		arg.AuthorName,
		arg.AverageRating,
		arg.AuthorID,
		arg.TextReviewCount,
		arg.RatingsCount,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.AuthorName,
		&i.AverageRating,
		&i.AuthorID,
		&i.TextReviewCount,
		&i.RatingsCount,
	)
	return i, err
}
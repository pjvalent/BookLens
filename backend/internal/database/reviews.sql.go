// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reviews.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createReview = `-- name: CreateReview :one
INSERT INTO reviews (id, created_at, updated_at, user_id, book_id, rating, review_text, spoiler_tag)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, updated_at, user_id, book_id, rating, review_text, spoiler_tag
`

type CreateReviewParams struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     uuid.UUID
	BookID     uuid.UUID
	Rating     int32
	ReviewText string
	SpoilerTag bool
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error) {
	row := q.db.QueryRowContext(ctx, createReview,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.BookID,
		arg.Rating,
		arg.ReviewText,
		arg.SpoilerTag,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.BookID,
		&i.Rating,
		&i.ReviewText,
		&i.SpoilerTag,
	)
	return i, err
}

const deleteReviewByID = `-- name: DeleteReviewByID :exec
DELETE FROM reviews WHERE user_id=$1 AND book_id=$2
`

type DeleteReviewByIDParams struct {
	UserID uuid.UUID
	BookID uuid.UUID
}

func (q *Queries) DeleteReviewByID(ctx context.Context, arg DeleteReviewByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteReviewByID, arg.UserID, arg.BookID)
	return err
}

const getAllUserReviews = `-- name: GetAllUserReviews :many
SELECT books.author, books.title, reviews.rating, reviews.review_text
FROM reviews JOIN books ON reviews.book_id = books.id
WHERE reviews.user_id = $1
`

type GetAllUserReviewsRow struct {
	Author     string
	Title      string
	Rating     int32
	ReviewText string
}

func (q *Queries) GetAllUserReviews(ctx context.Context, userID uuid.UUID) ([]GetAllUserReviewsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllUserReviews, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUserReviewsRow
	for rows.Next() {
		var i GetAllUserReviewsRow
		if err := rows.Scan(
			&i.Author,
			&i.Title,
			&i.Rating,
			&i.ReviewText,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReviewByUserIDBookID = `-- name: GetReviewByUserIDBookID :one
SELECT COUNT(*) FROM reviews WHERE user_id=$1 AND book_id=$2
`

type GetReviewByUserIDBookIDParams struct {
	UserID uuid.UUID
	BookID uuid.UUID
}

func (q *Queries) GetReviewByUserIDBookID(ctx context.Context, arg GetReviewByUserIDBookIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getReviewByUserIDBookID, arg.UserID, arg.BookID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateReview = `-- name: UpdateReview :one
UPDATE reviews
SET 
    rating = $1,
    review_text = $2,
    updated_at = $3
WHERE 
    user_id = $4 AND book_id = $5
RETURNING id, created_at, updated_at, user_id, book_id, rating, review_text, spoiler_tag
`

type UpdateReviewParams struct {
	Rating     int32
	ReviewText string
	UpdatedAt  time.Time
	UserID     uuid.UUID
	BookID     uuid.UUID
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) (Review, error) {
	row := q.db.QueryRowContext(ctx, updateReview,
		arg.Rating,
		arg.ReviewText,
		arg.UpdatedAt,
		arg.UserID,
		arg.BookID,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.BookID,
		&i.Rating,
		&i.ReviewText,
		&i.SpoilerTag,
	)
	return i, err
}

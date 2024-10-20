// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: books.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (id, isbn, created_at, updated_at, title, author, num_pages, price)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, isbn, title, author, num_pages, price, created_at, updated_at
`

type CreateBookParams struct {
	ID        uuid.UUID
	Isbn      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Author    string
	NumPages  int32
	Price     int32
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.ID,
		arg.Isbn,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Author,
		arg.NumPages,
		arg.Price,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Isbn,
		&i.Title,
		&i.Author,
		&i.NumPages,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBookByTitleAuthor = `-- name: GetBookByTitleAuthor :one
SELECT id, isbn, title, author, num_pages, price, created_at, updated_at FROM books WHERE title=$1 AND author=$2
`

type GetBookByTitleAuthorParams struct {
	Title  string
	Author string
}

func (q *Queries) GetBookByTitleAuthor(ctx context.Context, arg GetBookByTitleAuthorParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBookByTitleAuthor, arg.Title, arg.Author)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Isbn,
		&i.Title,
		&i.Author,
		&i.NumPages,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

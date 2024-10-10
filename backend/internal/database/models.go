// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID
	Isbn      string
	Title     string
	Author    string
	NumPages  int32
	Price     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BooksGenere struct {
	Isbn     string
	GenereID uuid.UUID
	ID       uuid.NullUUID
}

type Genere struct {
	GenereID   uuid.UUID
	GenereName string
}

type Review struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     uuid.UUID
	BookID     uuid.UUID
	Rating     int32
	ReviewText string
	SpoilerTag bool
}

type User struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	FirstName      string
	LastName       string
	Email          string
	AccountBalance int64
	ApiKey         string
}
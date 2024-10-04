// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
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
	ID       int32
	Isbn     string
	GenereID int32
}

type Genere struct {
	GenereID int32
	Name     sql.NullString
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

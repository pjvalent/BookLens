package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pjvalent/BookLens/backend/internal/database"
)

type Book struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Isbn      string    `json:"isbn"`
	Author    string    `json:"author"`
	NumPages  int32     `json:"num_pages"`
	Price     int32     `json:"price"`
	Publisher string    `json:"publisher"`
	BookDesc  string    `json:"book_desc"`
}

func ConvertDbBookToBook(dbBook database.Book) Book {
	return Book{
		ID:        dbBook.ID,
		CreatedAt: dbBook.CreatedAt,
		UpdatedAt: dbBook.UpdatedAt,
		Isbn:      dbBook.Isbn,
		Author:    dbBook.Author,
		NumPages:  dbBook.NumPages,
		Price:     dbBook.Price,
		Publisher: dbBook.Publisher.String,
		BookDesc:  dbBook.BookDesc.String,
	}
}

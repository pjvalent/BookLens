package models

import (
	"github.com/pjvalent/BookLens/backend/internal/database"
)

type Reccomendation struct {
	Title     string
	Publisher string
	BookDesc  string
}

func ConvertDbBookRecToBookRec(dbSimilarBookByDescRow []database.SimilarBooksByDescRow) []Reccomendation {
	recommendedBooks := make([]Reccomendation, len(dbSimilarBookByDescRow))

	for i, book := range dbSimilarBookByDescRow {
		newBook := Reccomendation{
			Title:     book.Title,
			Publisher: book.Publisher.String,
			BookDesc:  book.BookDesc.String,
		}
		recommendedBooks[i] = newBook
	}

	return recommendedBooks
}

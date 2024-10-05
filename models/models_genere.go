package models

import (
	"github.com/google/uuid"
	"github.com/pjvalent/BookLens/internal/database"
)

type Genere struct {
	GenereID   uuid.UUID `json:"genere_id"`
	GenereName string    `json:"genere_name"`
}

func ConvertDbGenereToGenere(dbGenere database.Genere) Genere {
	return Genere{
		GenereID:   dbGenere.GenereID,
		GenereName: dbGenere.GenereName,
	}
}

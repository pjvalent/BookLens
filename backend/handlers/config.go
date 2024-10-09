package handlers

import "github.com/pjvalent/BookLens/backend/internal/database"

type ApiConfig struct {
	DB *database.Queries
}

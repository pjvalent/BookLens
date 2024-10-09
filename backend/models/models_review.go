package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pjvalent/BookLens/backend/internal/database"
)

type Review struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserID     uuid.UUID `json:"user_id"`
	BookID     uuid.UUID `json:"book_id"`
	Rating     int32     `json:"rating"`
	ReviewText string    `json:"review_text"`
	SpoilerTag bool      `json:"spoiler_tag"`
}

func ConvertDbReviewToReview(dbReview database.Review) Review {
	return Review{
		ID:         dbReview.ID,
		CreatedAt:  dbReview.CreatedAt,
		UpdatedAt:  dbReview.UpdatedAt,
		UserID:     dbReview.UserID,
		BookID:     dbReview.BookID,
		Rating:     dbReview.Rating,
		ReviewText: dbReview.ReviewText,
		SpoilerTag: dbReview.SpoilerTag,
	}
}

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

type UserReview struct {
	Author     string `json:"author"`
	Title      string `json:"title"`
	Rating     int32  `json:"rating"`
	ReviewText string `json:"review_text"`
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

func ConvertDbUserReviewToUserReview(dbUserReview database.GetAllUserReviewsRow) UserReview {
	return UserReview{
		Author:     dbUserReview.Author,
		Title:      dbUserReview.Title,
		Rating:     dbUserReview.Rating,
		ReviewText: dbUserReview.ReviewText,
	}
}

func ConvertDbUserReviewListToUserReviewList(dbUserReviewList []database.GetAllUserReviewsRow) []UserReview {
	userReviews := make([]UserReview, len(dbUserReviewList))

	for i, review := range dbUserReviewList {
		newReview := UserReview{
			Author:     review.Author,
			Title:      review.Title,
			Rating:     review.Rating,
			ReviewText: review.ReviewText,
		}
		userReviews[i] = newReview
	}

	return userReviews
}

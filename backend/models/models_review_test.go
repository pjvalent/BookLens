package models

import (
	"reflect"
	"testing"

	"github.com/pjvalent/BookLens/backend/internal/database"
)

func TestConvertDbUserReviewListToUserReviewList(t *testing.T) {
	tests := []struct {
		name                string
		input               []database.GetAllUserReviewsRow
		expectedUserReviews []UserReview
	}{
		{
			name: "Basic Conversion",
			input: []database.GetAllUserReviewsRow{
				{
					Author:     "George Orwell",
					Title:      "1984",
					Rating:     5,
					ReviewText: "A timeless classic about the dangers of totalitarianism.",
				},
				{
					Author:     "J.K. Rowling",
					Title:      "Harry Potter and the Sorcerer's Stone",
					Rating:     4,
					ReviewText: "An enchanting start to a magical series.",
				},
			},
			expectedUserReviews: []UserReview{
				{
					Author:     "George Orwell",
					Title:      "1984",
					Rating:     5,
					ReviewText: "A timeless classic about the dangers of totalitarianism.",
				},
				{
					Author:     "J.K. Rowling",
					Title:      "Harry Potter and the Sorcerer's Stone",
					Rating:     4,
					ReviewText: "An enchanting start to a magical series.",
				},
			},
		},
		{
			name:                "Empty Input",
			input:               []database.GetAllUserReviewsRow{},
			expectedUserReviews: []UserReview{},
		},
		{
			name: "Special Characters",
			input: []database.GetAllUserReviewsRow{
				{
					Author:     "Gabriel García Márquez",
					Title:      "Cien años de soledad",
					Rating:     5,
					ReviewText: "Una obra maestra de la literatura latinoamericana.",
				},
			},
			expectedUserReviews: []UserReview{
				{
					Author:     "Gabriel García Márquez",
					Title:      "Cien años de soledad",
					Rating:     5,
					ReviewText: "Una obra maestra de la literatura latinoamericana.",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userReviews := ConvertDbUserReviewListToUserReviewList(tt.input)
			if !reflect.DeepEqual(userReviews, tt.expectedUserReviews) {
				t.Errorf("Expected %v, got %v", tt.expectedUserReviews, userReviews)
			}
		})
	}
}

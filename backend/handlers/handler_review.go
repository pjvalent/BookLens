package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pjvalent/BookLens/backend/internal/database"
	"github.com/pjvalent/BookLens/backend/models"
)

func (apiCfg *ApiConfig) HandlerCreateReview(w http.ResponseWriter, r *http.Request, user database.User) {

	// id, created_at, updated_at are all generated at time of call
	// user_id, book_id will be looked up
	// as this is an authenticated endpoint, user_id will be fine
	// book_id will need to be looked up as we take title/author at endpoint, if book doesn't exist, return an error
	type parameters struct {
		Title      string `json:"title"`
		Author     string `json:"author"`
		Rating     int32  `json:"rating"`
		ReviewText string `json:"review_text"`
		SpoilerTag bool   `json:"spoiler_tag"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding user while creating review: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Check to see if the book by title/author exists, if it does not respond with error
	book, err := apiCfg.DB.GetBookByTitleAuthor(r.Context(), database.GetBookByTitleAuthorParams{
		Title:  params.Title,
		Author: params.Author,
	})

	if err != nil {
		log.Printf("Error creating review for book Title: %v --- Author: %v |||| %v", params.Title, params.Author, err)
		RespondWithError(w, http.StatusNoContent, fmt.Sprintf("Could not find book with provided title/author: %v", err))
		return
	}

	// if it does exist, see if the user has already created a review for it in the reviews table.
	count, err := apiCfg.DB.GetReviewByUserIDBookID(r.Context(), database.GetReviewByUserIDBookIDParams{
		UserID: user.ID,
		BookID: book.ID,
	})

	if err != nil {
		log.Printf("Error in getReviewByUsserIDBookID")
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Issue with creating review %v", err))
		return
	}

	if count != 0 {
		log.Printf("User trying to review book they already reviewed: %v", book.Title)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Already reviewed book: %v", params.Title))
		return
	}

	//TODO: need to check if the user has reviewed the book already, if they have need to be made to update review instead of create

	// check if the review text is empty, store as bool, check
	review, err := apiCfg.DB.CreateReview(r.Context(), database.CreateReviewParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		UserID:     user.ID,
		BookID:     book.ID,
		Rating:     params.Rating,
		ReviewText: params.ReviewText,
		SpoilerTag: params.SpoilerTag,
	})

	if err != nil {
		log.Printf("Error creating review: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error creating review: %v", err))
		return
	}

	RespondWithJSON(w, 201, models.ConvertDbReviewToReview(review))
}

func (apiCfg *ApiConfig) HandlerGetUserReviews(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userIDContextKey).(string)

	if !ok {
		log.Printf("Error gettting userIDContextKey: %v", ok)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting user reviews: %v", ok))
		return
	}

	userIDAsUUID, err := uuid.Parse(userID)
	if err != nil {
		log.Printf("Error with uuid: %v", err)
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user reviews: %v", err))
		return
	}

	reviews, err := apiCfg.DB.GetAllUserReviews(r.Context(), userIDAsUUID)

	if err != nil {
		log.Printf("Error with getting all user reviews: %v", err)
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user reviews: %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, models.ConvertDbUserReviewListToUserReviewList(reviews))

}

func (apiCfg *ApiConfig) HandlerDeleteReview(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		BookID string `json:"book_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding user while creating review: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	userID, ok := r.Context().Value(userIDContextKey).(string)

	if !ok {
		log.Printf("Error parsing the userID from the context")
		RespondWithError(w, 500, fmt.Sprintf("Internal server error :("))
		return
	}

	userIDAsUUID, err := uuid.Parse(userID)
	if err != nil {
		log.Printf("Error decoding user while deleting review: %v", err)
		RespondWithError(w, 500, fmt.Sprintf("Error parsing user: %v", err))
		return
	}

	bookIdAsUUID, err := uuid.Parse(params.BookID)

	if err != nil {
		log.Printf("Error decoding book id: %v", err)
		RespondWithError(w, 500, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	err = apiCfg.DB.DeleteReviewByID(r.Context(), database.DeleteReviewByIDParams{
		UserID: userIDAsUUID,
		BookID: bookIdAsUUID,
	})

	if err != nil {
		log.Printf("Error decoding user while deleting review: %v", err)
		RespondWithError(w, 500, fmt.Sprintf("Error deleting review: %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "success",
	})

}

// TODO: figure  out what the heck I was doing here
// func (apiCfg *ApiConfig) HandlerDeleteReview(w http.ResponseWriter, r *http.Request) {
// 	userID, ok := r.Context().Value(userIDContextKey).(string)

// 	if !ok {
// 		log.Printf("Error parsing the userID from the context")
// 		RespondWithError(w, 500, fmt.Sprintf("Internal server error :("))
// 		return
// 	}

// }

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pjvalent/BookLens/backend/internal/database"
	"github.com/pjvalent/BookLens/backend/internal/security"
	"github.com/pjvalent/BookLens/backend/internal/validate"
	"github.com/pjvalent/BookLens/backend/models"
)

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	// TODO: Update so that account balance can be a dollar value, then convert to cents for storing in the database

	type parameters struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		// AccountBalance int64  `json:"account_balance"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding user while creating user: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// check for valid email
	err = validate.ValidateEmail(params.Email)
	if err != nil {
		log.Printf("Error parsing email: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing email: %v", err))
		return
	}

	password, err := security.HashPassword(params.Password)

	if err != nil {
		log.Printf("Error with creating user password: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Password error: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		FirstName:      params.FirstName,
		LastName:       params.LastName,
		Email:          params.Email,
		AccountBalance: 0,
		UserPassword:   string(password),
	})

	if err != nil {
		log.Printf("Error creating user: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	RespondWithJSON(w, 201, models.ConvertDbUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, 200, models.ConvertDbUserToUser(user))

}

func (apiCfg *ApiConfig) HandlerGetUserByToken(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userIDContextKey).(string)

	if !ok {
		log.Printf("Error, issue getting context value userIDContextKey")
		RespondWithError(w, 500, fmt.Sprintf("Error getting context value userIDContextKey: %v", ok))
		return
	}

	userIDAsUUID, err := uuid.Parse(userID)

	if err != nil {
		log.Printf("Error, issue parsing userID")
		RespondWithError(w, 500, fmt.Sprintf("Error issue parsing userID: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByID(r.Context(), userIDAsUUID)

	if err != nil {
		log.Printf("Error with GetUserByID")
		RespondWithError(w, 500, fmt.Sprintf("Error fetching user: %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, models.ConvertDbUserToUser(user))

}

func (apiCfg *ApiConfig) HandlerGetAllUserReviews(w http.ResponseWriter, r *http.Request, user database.User) {

	reviews, err := apiCfg.DB.GetAllUserReviews(r.Context(), user.ID)

	if err != nil {
		log.Printf("Error fetching user reviews: %v", err)
		RespondWithError(w, 500, fmt.Sprintf("Error fetching user reviews: %v", err))
		return
	}

	userReviews := make([]models.UserReview, len(reviews))

	for i, dbUserReview := range reviews {
		userReviews[i] = models.ConvertDbUserReviewToUserReview(dbUserReview)
	}

	RespondWithJSON(w, 200, userReviews)

}

func (apiConfig *ApiConfig) HandlerDeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	err := apiConfig.DB.DeleteUserByUserID(r.Context(), user.ID)

	if err != nil {
		log.Printf("Error deleting user: %v", err)
		RespondWithError(w, 500, fmt.Sprintf("Error deleting user: %v", err))
		return
	}

	RespondWithJSON(w, 200, struct {
		Status string `json:"status"`
	}{
		Status: "success",
	})

}

func (apiCfg *ApiConfig) HandlerDeleteUserV1(w http.ResponseWriter, r *http.Request) {
	// get the userID
	userID, ok := r.Context().Value(userIDContextKey).(string)

	if !ok {
		log.Printf("Error deleting user, issue getting context value userIDContextKey")
		RespondWithError(w, 500, fmt.Sprintf("Error deleting user"))
		return
	}
	userIDAsUUID, err := uuid.Parse(userID)
	if err != nil {
		log.Printf("unable to parse uuid")
		RespondWithError(w, 500, fmt.Sprintf("Error deleting user"))
		return
	}

	err = apiCfg.DB.DeleteUserByUserID(r.Context(), userIDAsUUID)
	if err != nil {
		log.Printf("unable to delete user")
		RespondWithError(w, 500, fmt.Sprintf("Error deleting user"))
		return
	}

	RespondWithJSON(w, 200, struct {
		Status string `json:"status"`
	}{
		Status: "success",
	})
}

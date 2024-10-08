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

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	// TODO: Update so that account balance can be a dollar value, then convert to cents for storing in the database

	type parameters struct {
		FirstName      string `json:"first_name"`
		LastName       string `json:"last_name"`
		Email          string `json:"email"`
		AccountBalance int64  `json:"account_balance"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding user while creating user: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		FirstName:      params.FirstName,
		LastName:       params.LastName,
		Email:          params.Email,
		AccountBalance: params.AccountBalance,
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

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pjvalent/BookLens/internal/database"
	"github.com/pjvalent/BookLens/models"
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

	RespondWithJSON(w, 200, models.ConvertDbUserToUser(user))
}

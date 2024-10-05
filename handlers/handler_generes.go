package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/pjvalent/BookLens/internal/database"
	"github.com/pjvalent/BookLens/models"
)

func (apiCfg *ApiConfig) HandlerCreateGenere(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		GenereName string `json:"genere_name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding user while creating user: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	genere, err := apiCfg.DB.CreateGenere(r.Context(), database.CreateGenereParams{
		GenereID:   uuid.New(),
		GenereName: params.GenereName,
	})

	if err != nil {
		log.Printf("Error creating book: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error creating genere: %v", err))
		return
	}

	RespondWithJSON(w, 201, models.ConvertDbGenereToGenere(genere))
}

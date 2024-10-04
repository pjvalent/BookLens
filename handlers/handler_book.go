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

func (apiCfg *ApiConfig) HandlerCreateBook(w http.ResponseWriter, r *http.Request, user database.User) {

	// TODO: Update so that account balance can be a dollar value, then convert to cents for storing in the database

	type parameters struct {
		Isbn     string `json:"isbn"`
		Title    string `json:"title"`
		Author   string `json:"author"`
		NumPages int32  `json:"num_pages"`
		Price    int32  `json:"price"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding user while creating user: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	book, err := apiCfg.DB.CreateBook(r.Context(), database.CreateBookParams{
		ID:        uuid.New(),
		Isbn:      params.Isbn,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     params.Title,
		Author:    params.Author,
		NumPages:  params.NumPages,
		Price:     params.Price,
	})

	if err != nil {
		log.Printf("Error creating book: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error creating book: %v", err))
		return
	}

	RespondWithJSON(w, 201, models.ConvertDbBookToBook(book))
}

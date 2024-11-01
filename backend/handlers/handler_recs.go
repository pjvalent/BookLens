package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pjvalent/BookLens/backend/internal/database"
	"github.com/pjvalent/BookLens/backend/models"
)

func (apiCfg *ApiConfig) HandlerGetSimilarBooksByDesc(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Isbn  string `json:"isbn"`
		Title string `json:"title"`
		Limit int32  `json:"limit"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding params while creating getting similar books: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Take the isbn, do a search on the embeddings to find the closest ones

	similarBooks, err := apiCfg.DB.SimilarBooksByDesc(r.Context(), database.SimilarBooksByDescParams{
		Isbn:  params.Isbn,
		Limit: params.Limit,
	})

	if err != nil {
		log.Printf("Error getting similar books: %v", err)
		RespondWithError(w, 500, fmt.Sprintf("Error fetching similar books: %v", err))
		return
	}

	RespondWithJSON(w, 200, models.ConvertDbBookRecToBookRec(similarBooks))

}

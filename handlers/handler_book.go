package handlers

import (
	"database/sql"
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
		Isbn     string   `json:"isbn"`
		Title    string   `json:"title"`
		Author   string   `json:"author"`
		NumPages int32    `json:"num_pages"`
		Price    int32    `json:"price"`
		Generes  []string `json:"generes"`
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

	//TODO: Parse the generes from the Genere field in the json
	//		After parsing, lookup each genere to see if it exists
	//		For each genere, attempt to fetch the genere from the database
	//		If it doesn't exist, create a new genere in the generes table
	//		After you have the genere, insert a record into books_generes with the correct book,genere

	for _, genereName := range params.Generes {
		genere, err := apiCfg.DB.GetGenereByName(r.Context(), genereName)

		if err != nil {
			//check to see if it is a sql.ErrNoRows error, if it is then the genere doesn't exist and we should make a new one
			if err == sql.ErrNoRows {
				genere, err = apiCfg.DB.CreateGenere(r.Context(), database.CreateGenereParams{
					GenereID:   uuid.New(),
					GenereName: genereName,
				})
				if err != nil {
					log.Printf("Error creating genre '%s': %v", genereName, err)
					RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating genre '%s'", genereName))
					return
				}
			} else {
				log.Printf("Error fetching genre '%s': %v", genereName, err)
				RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching genre '%s'", genereName))
				return
			}
		}

		//err != nil so we found the genere
		_, err = apiCfg.DB.CreateBooksGeneres(r.Context(), database.CreateBooksGeneresParams{
			Isbn:     params.Isbn,
			GenereID: genere.GenereID,
		})

		if err != nil {
			log.Printf("Error creating booksGenere entry %v", err)
			RespondWithError(w, 400, fmt.Sprintf("Error creating booksGenere entry: %v", err))
		}
	}

	RespondWithJSON(w, 201, models.ConvertDbBookToBook(book))
}

package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pjvalent/BookLens/backend/internal/auth"
	"github.com/pjvalent/BookLens/backend/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// TODO: add password checking to this function
func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			log.Printf("error authenticating api key: %v", err)
			RespondWithError(w, 403, fmt.Sprintf("Error authenticating api key: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			log.Printf("error querying user by api key: %v", err)
			RespondWithError(w, 400, fmt.Sprintf("Error, unable to find user: %v", err))
			return
		}

		handler(w, r, user)

	}
}

// // TODO: User will login on a login handler endpoint, that endpoint will generate a jwt token for the user and pass it back to the user
// // this middleware will be an auth middleware that validates with the jwt token instead of the api token
// func (cfg *ApiConfig) UserPasswordAuth(handler authedHandler) http.HandlerFunc {

// }

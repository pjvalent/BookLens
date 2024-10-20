package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/pjvalent/BookLens/backend/internal/auth"
	"github.com/pjvalent/BookLens/backend/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)
type tokenHandler func(http.ResponseWriter, *http.Request)

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

// // TODO: Change to utilize token in cookie
func (apiCfg *ApiConfig) MiddlewareTokenAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")

		if err != nil {
			log.Printf("error parsing cookie: %v", err)
			RespondWithError(w, 403, fmt.Sprintf("Cookie error: %v", err))
			return
		}

		tokenString := cookie.Value

		// tokenString, err := auth.GetToken(r.Header)

		// if err != nil {
		// 	log.Printf("error authenticating bearer token: %v", err)
		// 	RespondWithError(w, 403, fmt.Sprintf("Error authenticating bearer token: %v", err))
		// 	return
		// }

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return apiCfg.JWTKey, nil
		})

		if err != nil || !token.Valid {
			RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Attach user ID to the request context
		ctx := context.WithValue(r.Context(), userIDContextKey, claims.UserID)
		r = r.WithContext(ctx)

		handler(w, r)
	}
}

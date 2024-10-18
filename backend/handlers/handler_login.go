package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pjvalent/BookLens/backend/internal/security"
)

func (apiCfg *ApiConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding user while creating user: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)

	if err != nil {
		log.Printf("Error can't fine user by email: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error, can't find user with provided emaiol: %v", err))
		return
	}

	// check if the hashed password from the database matches the provided password
	err = security.CheckPassword(user.UserPassword, params.Password)

	if err != nil {
		log.Printf("Error password does not match: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error, incorrect password: %v", err))
		return
	}

	// if no err, passwords match and now we generate and respond with a jwt token
	token, err := apiCfg.GenerateToken(user.ID.String())

	if err != nil {
		log.Printf("Error generating token: %v", err)
		RespondWithError(w, 500, fmt.Sprintf("Error generating token: %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})

}

// TODO: get this out of the login handler, put it in its own file, also make a helpers/utilities directory to put this function/claims/json/config in it
func (apiCfg *ApiConfig) GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "BookLense",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(apiCfg.JWTKey)
}

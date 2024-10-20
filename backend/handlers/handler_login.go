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
		log.Printf("Error can't find user by email: %v", err)
		RespondWithError(w, 400, fmt.Sprintf("Error, can't find user with provided email: %v", err))
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

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"status": "success",
	})

}

func (apiCfg *ApiConfig) HandlerLogout(w http.ResponseWriter, r *http.Request) {
	// Clear the token cookie to make the user be logged out
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	RespondWithJSON(w, http.StatusOK, map[string]string{"status": "logged out"})
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

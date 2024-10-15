package handlers

import "github.com/golang-jwt/jwt/v5"

type contextKey string

const userIDContextKey contextKey = "userID"

// Claims defines the structure of the JWT claims.
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

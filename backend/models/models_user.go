package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pjvalent/BookLens/backend/internal/database"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	AccountBalance int64     `json:"account_balance"`
	ApiKey         string    `json:"api_key"`
	UserPassword   string    `json:"password"`
}

func ConvertDbUserToUser(dbUser database.User) User {
	return User{
		ID:             dbUser.ID,
		CreatedAt:      dbUser.CreatedAt,
		UpdatedAt:      dbUser.UpdatedAt,
		FirstName:      dbUser.FirstName,
		LastName:       dbUser.LastName,
		Email:          dbUser.Email,
		AccountBalance: dbUser.AccountBalance,
		ApiKey:         dbUser.ApiKey,
		UserPassword:   dbUser.UserPassword,
	}
}

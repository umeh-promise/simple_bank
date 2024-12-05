// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"
)

type Account struct {
	ID          int64     `json:"id"`
	Owner       string    `json:"owner"`
	Balance     int64     `json:"balance"`
	Currency    string    `json:"currency"`
	CountryCode int32     `json:"country_code"`
	CreatedAt   time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// It might be positive or negative
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// It must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Username        string    `json:"username"`
	HashedPassword  string    `json:"hashed_password"`
	FullName        string    `json:"full_name"`
	Email           string    `json:"email"`
	PasswordChanged time.Time `json:"password_changed"`
	CreatedAt       time.Time `json:"created_at"`
}

package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserID int64

type User struct {
	ID UserID `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
	Role string `json:"role" db:"role"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Modified_at time.Time `json:"modified_at" db:"modified_at"`
}

func (u *User) ComparePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(pw))
}
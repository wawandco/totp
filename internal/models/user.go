package models

import (
	"database/sql"
	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID       uuid.UUID      `json:"id" db:"id"`
	Email    string         `json:"email" db:"email"`
	Password sql.NullString `json:"password" db:"password"`
	Secret   sql.NullString `json:"secret" db:"secret"`
}

type UsersService interface {
	Find(email string) (User, error)
	SetSecret(email string, secret string) error
}

package users

import (
	"database/sql"
	"github.com/dmartinez24/totp/internal/models"
)

type service struct {
	db *sql.DB
}

func NewService(db *sql.DB) models.UsersService {
	return &service{db: db}
}

func (s *service) Find(email string) (user models.User, err error) {
	err = s.db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Secret, &user.Password)
	if err != nil {
		return user, err
	}

	return user, err
}

func (s *service) SetSecret(email string, secret string) error {
	_, err := s.db.Exec("UPDATE  users SET secret = ? WHERE email = ?", secret, email)

	if err != nil {
		return err
	}

	return nil
}

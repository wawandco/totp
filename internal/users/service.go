package users

import (
	"database/sql"
	"easytotp/internal/models"
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
		return
	}

	return
}

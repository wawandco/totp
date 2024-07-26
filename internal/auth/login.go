package auth

import (
	"facts/internal/models"
	"fmt"
	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server/session"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
)

func Validate(w http.ResponseWriter, r *http.Request) {
	s := session.FromCtx(r.Context())
	usersService := r.Context().Value("users").(models.UsersService)
	rw := render.FromCtx(r.Context())

	email := r.FormValue("email")
	password := r.FormValue("password")

	showError := func(log string, err error) {
		slog.Error(fmt.Sprintf("%s: %v", log, err))
		rw.Set("toast", "❌ invalid email or password")
		err = rw.Render("auth/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	user, err := usersService.Find(email)

	if err != nil {
		showError("error loading user", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))

	if err != nil {
		showError("error comparing hash and password", err)
		return
	}

	s.Values[FirstDoorKey] = email

	err = s.Save(r, w)

	if err != nil {
		slog.Error(fmt.Sprintf("error saving session %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/authenticate", http.StatusSeeOther)
}

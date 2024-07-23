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

	user, err := usersService.Find(email)

	if err != nil {
		slog.Error(fmt.Sprintf("error loading user: %v", err))
		s.AddFlash("❌ invalid email or password", "toasts")
		rw.Set("toasts", s.Flashes("toasts"))
		err = rw.Render("auth/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))

	if err != nil {
		slog.Error(fmt.Sprintf("error comparing hash and password: %v", err))
		s.AddFlash("❌ invalid email or password", "toasts")
		rw.Set("toasts", s.Flashes("toasts"))
		err = rw.Render("auth/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	s.Values[SessionKey] = email

	err = s.Save(r, w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

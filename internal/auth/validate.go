package auth

import (
	"fmt"
	"github.com/dmartinez24/totp/internal/models"
	"github.com/dmartinez24/totp/internal/totp"
	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server/session"
	"log/slog"
	"net/http"
	"strings"
)

func Validate(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	s := session.FromCtx(r.Context())
	user := r.Context().Value("currentUser").(models.User)
	authenticator := r.Context().Value("totp").(totp.Authenticator)
	
	err := r.ParseForm()
	if err != nil {
		slog.Error(fmt.Sprintf("error parsing form: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	template := "auth/setup.html"
	secret := r.FormValue("secret")
	qr := r.FormValue("qr")
	code := strings.Join(r.Form["code"], "")

	if user.Secret.Valid {
		secret = user.Secret.String
		template = "auth/validate.html"
	}

	valid := authenticator.Validate(code, secret)

	if !valid {
		rw.Set("qr", qr)
		rw.Set("secret", secret)
		rw.Set("toast", "❌ invalid code")
		err = rw.Render(template)
		if err != nil {
			slog.Error(fmt.Sprintf("error rendering template: %v", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if !user.Secret.Valid && secret != "" {
		usersService := r.Context().Value("users").(models.UsersService)
		err = usersService.SetSecret(user.Email, secret)

		if err != nil {
			slog.Error(fmt.Sprintf("error updating user: %v", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	s.Values[SecondDoorKey] = true
	err = s.Save(r, w)

	if err != nil {
		slog.Error(fmt.Sprintf("error saving session: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

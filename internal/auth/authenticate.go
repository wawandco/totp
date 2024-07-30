package auth

import (
	"easytotp/internal/models"
	"easytotp/internal/totp"
	"fmt"

	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server/session"
	"log/slog"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	s := session.FromCtx(r.Context())

	if s.Values[SecondDoorKey] != nil {
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}

	user, ok := r.Context().Value("currentUser").(models.User)

	if !ok {
		slog.Info("no user in context")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	template := "auth/verify.html"

	if !user.Secret.Valid {
		authenticator := r.Context().Value("totp").(totp.Authenticator)

		secret, err := authenticator.GenerateSecretKey(user.Email)
		if err != nil {
			slog.Error(fmt.Sprintf("error generating key: %v", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		qr, err := authenticator.GenerateQR()
		if err != nil {
			slog.Error(fmt.Sprintf("error generating qr: %v", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		template = "auth/setup.html"
		rw.Set("secret", secret)
		rw.Set("qr", qr)
	}

	err := rw.Render(template)
	if err != nil {
		slog.Error(fmt.Sprintf("error rendering template: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

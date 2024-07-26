package auth

import (
	"facts/internal/models"
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
	totpInfo, err := TOTPInfo(user.Secret.String, user.Email)

	if err != nil {
		slog.Error(fmt.Sprintf("error generating authenticator: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !user.Secret.Valid {
		template = "auth/setup.html"
		rw.Set("secret", totpInfo.Secret)
		rw.Set("qr", totpInfo.QR)
	}

	err = rw.Render(template)
	if err != nil {
		slog.Error(fmt.Sprintf("error rendering template: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package auth

import (
	"github.com/leapkit/leapkit/core/server/session"
	"net/http"

	"github.com/leapkit/leapkit/core/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	s := session.FromCtx(r.Context())

	if s.Values[SessionKey] != nil {
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}

	rw := render.FromCtx(r.Context())

	err := rw.Render("auth/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

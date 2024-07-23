package auth

import (
	"fmt"
	"github.com/leapkit/leapkit/core/server/session"
	"log/slog"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	s := session.FromCtx(r.Context())

	for k := range s.Values {
		s.Values[k] = nil
	}

	if err := s.Save(r, w); err != nil {
		slog.Error(fmt.Sprintf("error saving the session: %v\n", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

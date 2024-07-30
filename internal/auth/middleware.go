package auth

import (
	"easytotp/internal/models"
	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server/session"
	"net/http"
)

const SessionKey = "USER_KEY"

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s := session.FromCtx(r.Context())
			rw := render.FromCtx(r.Context())

			if s.Values[SessionKey] == nil {
				err := rw.Render("auth/401.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			usersService := r.Context().Value("users").(models.UsersService)
			user, err := usersService.Find(s.Values[SessionKey].(string))

			if err != nil {
				err = rw.Render("auth/401.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			rw.Set("currentUser", user)

			next.ServeHTTP(w, r)
		})
	}
}

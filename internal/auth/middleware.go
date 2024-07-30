package auth

import (
	"context"
	"github.com/dmartinez24/totp/internal/models"

	"fmt"
	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server/session"
	"log/slog"
	"net/http"
)

const FirstDoorKey = "LOGIN"
const SecondDoorKey = "TOTP"

func LoginMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Info("LOGIN MW ---->")
			s := session.FromCtx(r.Context())
			rw := render.FromCtx(r.Context())

			if s == nil || s.Values[FirstDoorKey] == nil {
				err := rw.Render("auth/401.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			rw.Set("firstDoorOpen", true)

			slog.Info("<---- LOGIN MW")
			next.ServeHTTP(w, r)
		})
	}
}

func AuthenticatorMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Info("AUTHENTICATION MW ---->")
			s := session.FromCtx(r.Context())
			rw := render.FromCtx(r.Context())

			if s == nil || s.Values[SecondDoorKey] == nil {
				err := rw.Render("auth/401.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			rw.Set("secondDoorOpen", true)

			slog.Info("<---- AUTHENTICATION MW")
			next.ServeHTTP(w, r)
		})
	}
}

func CurrentUserMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Info("CURRENT USER MW ---->")
			s := session.FromCtx(r.Context())
			rw := render.FromCtx(r.Context())

			usersService := r.Context().Value("users").(models.UsersService)
			user, err := usersService.Find(s.Values[FirstDoorKey].(string))

			if err != nil {
				slog.Error(fmt.Sprintf("error loading user: %v", err))
				s.Values[FirstDoorKey] = nil
				s.Values[SecondDoorKey] = nil

				err = s.Save(r, w)
				if err != nil {
					slog.Error(fmt.Sprintf("error saving session: %v", err))
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				rw.Set("firstDoorOpen", false)
				rw.Set("secondDoorOpen", false)

				err = rw.Render("auth/401.html")
				if err != nil {
					slog.Error(fmt.Sprintf("error rendering template: %v", err))
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				slog.Info("<---- CURRENT USER MW")
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "currentUser", user)
			r = r.WithContext(ctx)

			rw.Set("currentUser", user)
			slog.Info("<---- CURRENT USER MW")
			next.ServeHTTP(w, r)
		})
	}
}

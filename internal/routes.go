package internal

import (
	"easytotp/internal/auth"
	"easytotp/internal/users"
	"easytotp/public"
	"github.com/leapkit/leapkit/core/assets"
	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server"
)

func AddRoutes(r server.Router) (err error) {
	assetsManager := assets.NewManager(public.Files)
	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),

		render.WithDefaultLayout("layout.html"),
		render.WithHelpers(render.AllHelpers),
		render.WithHelpers(map[string]any{
			"assetPath": assetsManager.PathFor,
		}),
	))

	r.HandleFunc("GET /login", auth.Index)
	r.HandleFunc("POST /login", auth.Login)
	r.HandleFunc("POST /logout", auth.Logout)

	r.Group("/", func(r server.Router) {
		r.Use(
			auth.LoginMiddleware(),
			auth.CurrentUserMiddleware(),
		)
		r.HandleFunc("GET /authenticate", auth.Authenticate)
		r.HandleFunc("POST /verify", auth.Verify)
	})

	r.Group("/", func(r server.Router) {
		r.Use(
			auth.LoginMiddleware(),
			auth.AuthenticatorMiddleware(),
			auth.CurrentUserMiddleware(),
		)
		r.HandleFunc("GET /{$}", users.Index)
		r.HandleFunc("GET /users", users.Index)
	})

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	r.HandleFunc(assetsManager.HandlerPattern(), assetsManager.HandlerFn)

	return
}

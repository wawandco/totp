package internal

import (
	"cmp"
	"embed"
	"fmt"
	"github.com/leapkit/leapkit/core/db"
	"github.com/leapkit/leapkit/core/server"
	"net/http"
	"os"
)

var (
	//go:embed **/*.html **/*.html *.html
	tmpls embed.FS

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(
		cmp.Or(os.Getenv("DATABASE_URL"), "leapkit.db"),
		db.WithDriver("sqlite3"),
	)
)

// Server interface exposes the methods
// needed to start the server in the cmd/app package
type Server interface {
	Addr() string
	Handler() http.Handler
}

func New() Server {
	// Creating a new server instance with the
	// default host and port values.
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3001")),
		server.WithSession(
			cmp.Or(os.Getenv("SESSION_SECRET"), "d720c059-9664-4980-8169-1158e167ae57"),
			cmp.Or(os.Getenv("SESSION_NAME"), "leapkit_session"),
		),
	)

	// Application services
	if err := AddServices(r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Application services
	if err := AddRoutes(r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return r
}

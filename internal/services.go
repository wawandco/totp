package internal

import (
	"fmt"
	"github.com/dmartinez24/totp/internal/users"
	"github.com/leapkit/leapkit/core/server"
)

// AddServices is a function that will be called by the server
// to inject services in the context.
func AddServices(r server.Router) error {
	db, err := DB()
	if err != nil {
		return fmt.Errorf("connecting to the database: %w", err)
	}

	// Services that will be injected in the context
	r.Use(server.InCtxMiddleware("users", users.NewService(db)))

	return nil
}

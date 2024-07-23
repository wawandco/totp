package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"facts/internal"

	// Load environment variables
	_ "github.com/leapkit/leapkit/core/tools/envload"
	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	s := internal.New()

	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println("error starting app:", err)
	}
}

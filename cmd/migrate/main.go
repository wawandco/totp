package main

import (
	"facts/internal"
	"facts/internal/migrations"
	"fmt"

	"github.com/leapkit/leapkit/core/db"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

// The migrate command is used to ship our application
// with the latest database schema migrator. which can be invoked
// by running `migrate`.
func main() {
	conn, err := internal.DB()
	if err != nil {
		fmt.Println("error connecting to the database: ", err)
	}

	err = db.RunMigrations(migrations.All, conn)
	if err != nil {
		fmt.Println("error running migrations: ", err)
	}
}

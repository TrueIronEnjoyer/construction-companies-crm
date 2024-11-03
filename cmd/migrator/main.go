package main

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	absPath, _ := filepath.Abs("migrations")
	migrator, err := migrate.New(
		"file://"+absPath,
		"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	defer migrator.Close()

	if err = migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migration to aply")
		} else {
			panic(err)
		}
	} else {
		fmt.Println("Migrations applied!!")
	}
}

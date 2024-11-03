package migrator

import (
	"construction-companies-crm/internal/config"
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ApplyMigrations(conf *config.Config) {
	absPath, _ := filepath.Abs(conf.MigrationsPath)

	fmt.Printf("postgres://%s:%s@%s:%s/%s?sslmode=disable/n",
		conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	migrator, err := migrate.New(
		"file://"+absPath,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName),
	)
	if err != nil {
		panic(err)
	}
	defer migrator.Close()

	if err = migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			slog.Info("No migration to aply")
		} else {
			panic(err)
		}
	} else {
		slog.Info("Migrations applied!!")
	}
}
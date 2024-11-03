package main

import (
	"construction-companies-crm/internal/config"
	"construction-companies-crm/internal/migrator"
)

func main() {
	conf := config.NewConfig()

	migrator.ApplyMigrations(&conf)
}

package main

import (
	"construction-companies-crm/internal/config"
	"construction-companies-crm/internal/migrator"
)

func main() {
	config.InitConfig()
	conf := config.GetConfig()

	migrator.ApplyMigrations(conf)
}

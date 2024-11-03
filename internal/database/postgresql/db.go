package postgresql

import (
	"construction-companies-crm/internal/config"
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func DbConnect() {
	database, err := sql.Open("postgres", config.GetConfig().DbConfig.ConnStr)
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	return db.Close()
}

package config

import (
	"flag"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost         string
	DbPort         string
	DbUser         string
	DbPassword     string
	DbName         string
	Host           string
	Port           string
	Env            string
	MigrationsPath string
}

func NewConfig() Config {
	var configPath string
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()

	confMap, err := godotenv.Read(configPath)
	if err != nil {
		panic(err)
	}

	conf := Config{
		DbHost:         confMap["db_host"],
		DbPort:         confMap["db_port"],
		DbUser:         confMap["db_user"],
		DbPassword:     confMap["db_password"],
		DbName:         confMap["db_name"],
		Host:           confMap["host"],
		Port:           confMap["port"],
		Env:            confMap["env"],
		MigrationsPath: confMap["migrations_path"],
	}

	return conf
}

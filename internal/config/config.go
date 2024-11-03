package config

import (
	"flag"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig       DbConfig
	HostConfig     HostConfig
	Env            string
	MigrationsPath string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type HostConfig struct {
	Host string
	Port string
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
		DbConfig: DbConfig{
			Host:     confMap["db_host"],
			Port:     confMap["db_port"],
			User:     confMap["db_user"],
			Password: confMap["db_password"],
			Name:     confMap["db_name"],
		},
		HostConfig: HostConfig{
			Host: confMap["host"],
			Port: confMap["port"],
		},
		Env:            confMap["env"],
		MigrationsPath: confMap["migrations_path"],
	}

	return conf
}

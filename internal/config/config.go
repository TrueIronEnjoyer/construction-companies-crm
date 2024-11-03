package config

import (
	"errors"
	"flag"
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig               DbConfig
	HostConfig             HostConfig
	Env                    string
	MigrationsPath         string
	JwtSecret              string
	TokenTimeoutHoursCount int
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	ConnStr  string
}

type HostConfig struct {
	Host string
	Port string
}

var conf Config

func InitConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()

	confMap, err := godotenv.Read(configPath)
	if err != nil {
		panic(err)
	}

	conf = Config{
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
		JwtSecret:      confMap["jwt_secret"],
	}

	conf.DbConfig.ConnStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.DbConfig.User, conf.DbConfig.Password, conf.DbConfig.Host, conf.DbConfig.Port, conf.DbConfig.Name)

	conf.TokenTimeoutHoursCount, err = strconv.Atoi(confMap["token_timeout_hours_count"])
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func GetConfig() *Config {
	return &conf
}

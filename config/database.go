package config

import "os"

type DBConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
	TimeZone string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		TimeZone: os.Getenv("TIMEZONE"),
	}
}

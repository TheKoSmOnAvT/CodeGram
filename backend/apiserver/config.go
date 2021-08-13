package apiserver

import database "backend/Database"

type Config struct {
	BingAddr string
	LogLevel string
	database *database.Config
}

func NewConfig() *Config {
	return &Config{
		BingAddr: ":4000",
		LogLevel: "Debug",
		database: database.NewConfig(),
	}
}

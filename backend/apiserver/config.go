package apiserver

import database "backend/Database"

type Config struct {
	BingAddr string `toml: "bind_addr"`
	LogLevel string `toml: "log_level"`
	database *database.Config
}

func NewConfig() *Config {
	return &Config{
		BingAddr: ":4000",
		LogLevel: "Debug",
		database: database.NewConfig(),
	}
}

package database

type Config struct {
	DataBaseURL string `toml: "database_url"`
}

func NewConfig() *Config {
	return &Config{
		DataBaseURL: "../db/Codegram.db",
	}
}

package database

type Config struct {
	DataBaseURL string
}

func NewConfig() *Config {
	return &Config{
		DataBaseURL: "../db/Codegram.db",
	}
}

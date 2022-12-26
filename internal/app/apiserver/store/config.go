package store

type Config struct {
	DatabaseURL      string `toml:"database_url"`
	DatabaseUser     string `toml:"database_user"`
	DatabasePassword string `toml:"database_password"`
	DatabaseDBname   string `toml:"database_name"`
}

func NewConfig() *Config {
	return &Config{}
}

package apiserver

type Config struct {
	BindAddr         string `toml:"bind_addr"`
	LogLevel         string `toml:"log_level"`
	DatabaseURL      string `toml:"database_url"`
	DatabaseDBName   string `toml:"database_name"`
	DatabaseUser     string `toml:"database_user"`
	DatabasePassword string `toml:"database_password"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

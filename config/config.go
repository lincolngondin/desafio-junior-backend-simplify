package config

type Config struct {
	DBDriverName     string
	DBDataSourceName string
}

func New() *Config {
	return &Config{
		DBDriverName:     "sqlite",
		DBDataSourceName: "scripts/simplify.sqlite",
	}
}

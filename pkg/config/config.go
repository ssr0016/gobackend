package config

type Config struct {
	Postgres   PostgresConfig
	Server     ServerConfig
	Pagination PaginationConfig
	Auth       AuthConfig
}

func FromEnv() (*Config, error) {
	cfg := &Config{}

	cfg.postgresConfig()
	cfg.serverConfig()
	cfg.paginationConfig()
	cfg.authConfig()

	return cfg, nil
}

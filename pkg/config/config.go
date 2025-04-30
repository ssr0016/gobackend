package config

type Config struct {
	Postgres   PostgresConfig
	Server     ServerConfig
	Pagination PaginationConfig
}

func FromEnv() (*Config, error) {
	cfg := &Config{}

	cfg.postgresConfig()
	cfg.serverConfig()
	cfg.paginationConfig()

	return cfg, nil
}

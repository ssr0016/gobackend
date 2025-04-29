package config

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
}

func FromEnv() (*Config, error) {
	cfg := &Config{}

	cfg.postgresConfig()
	cfg.serverConfig()

	return cfg, nil
}

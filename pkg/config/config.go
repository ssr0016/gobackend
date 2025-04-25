package config

type Config struct {
	Server ServerConfig
}

func FromEnv() (*Config, error) {
	cfg := &Config{}

	cfg.serverConfig()

	return cfg, nil
}

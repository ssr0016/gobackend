package config

import "backend/pkg/util/env"

const (
	defaultServerPort = "4010"
)

type ServerConfig struct {
	HTTPPort string
}

func (cfg *Config) serverConfig() {
	cfg.Server.HTTPPort = env.GetEnvAsString("SERVER_HTTP_PORT", defaultServerPort)
}

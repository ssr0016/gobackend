package config

import (
	"backend/pkg/util/env"
	"time"
)

const (
	defaultAccessTokenDuration = 15
)

type AuthConfig struct {
	AccessTokenDuration time.Duration
}

func (cfg *Config) authConfig() {
	tokenDuration, _ := env.GetEnvAsInt("AUTH_ACCESS_TOKEN_DURATION", defaultAccessTokenDuration)
	cfg.Auth.AccessTokenDuration = time.Duration(tokenDuration) * time.Minute
}

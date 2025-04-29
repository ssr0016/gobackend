package config

import (
	"backend/pkg/util/env"
	"fmt"
)

const (
	defaultDBHost            = "localhost"
	defaultDBPort            = "5432" // Default PostgreSQL port
	defaultDBUser            = "postgres"
	defaultDBPassword        = "secret"
	defaultDBName            = "webuildfutureDB"
	defaultDBApplicationName = "webuildfuture"
	defaultSSLMode           = "disable" // Change to "disable" if needed
)

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DB              string
	ApplicationName string
	SSLMode         string
}

func (p *PostgresConfig) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s application_name=%s binary_parameters=yes",
		p.Host, p.Port, p.User, p.Password, p.DB, p.SSLMode, p.ApplicationName,
	)
}

func (cfg *Config) postgresConfig() {
	cfg.Postgres.Host = env.GetEnvAsString("DATABASE_HOST", defaultDBHost)
	cfg.Postgres.Port = env.GetEnvAsString("DATABASE_PORT", defaultDBPort)
	cfg.Postgres.User = env.GetEnvAsString("DATABASE_USER", defaultDBUser)
	cfg.Postgres.Password = env.GetEnvAsString("DATABASE_PASSWORD", defaultDBPassword)
	cfg.Postgres.DB = env.GetEnvAsString("DATABASE_NAME", defaultDBName)
	cfg.Postgres.ApplicationName = env.GetEnvAsString("DATABASE_APPLICATION_NAME", defaultDBApplicationName)
	cfg.Postgres.SSLMode = env.GetEnvAsString("DATABASE_SSLMODE", defaultSSLMode)
}

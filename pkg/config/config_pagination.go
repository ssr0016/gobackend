package config

import "backend/pkg/util/env"

const (
	defaultPage    = 1
	defaultPerPage = 10
)

type PaginationConfig struct {
	Page    int
	PerPage int
}

func (cfg *Config) paginationConfig() {
	cfg.Pagination.Page, _ = env.GetEnvAsInt("PAGINATION_PAGE", defaultPage)
	cfg.Pagination.PerPage, _ = env.GetEnvAsInt("PAGINATION_PER_PAGE", defaultPerPage)
}

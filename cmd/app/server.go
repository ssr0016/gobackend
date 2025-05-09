package app

import (
	"backend/pkg/config"
	"backend/pkg/identity/accesscontrol/accesscontrolimpl"
	"backend/pkg/identity/auth/auth"
	"backend/pkg/identity/user/userimpl"
	"backend/pkg/infra/api/auth/authimpl"
	"backend/pkg/infra/api/middleware"
	"backend/pkg/infra/registry"
	"backend/pkg/infra/storage/postgres"
	"backend/pkg/migration"
	"backend/pkg/protocol"
	"context"
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type Server struct {
	postgresDB postgres.DB
	services   []registry.RunFunc
	log        *zap.Logger
}

func NewServer(isStandaloneMode bool) (*Server, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	postgresDB, err := postgres.New(migration.New(), cfg.Postgres.ConnectionString())
	if err != nil {
		return nil, err
	}

	authSvc, err := authimpl.New()
	if err != nil {
		return nil, err
	}

	contextHandler, err := middleware.New(authSvc)
	if err != nil {
		return nil, err
	}
	accesscontrolSvc := accesscontrolimpl.NewService(postgresDB, cfg)
	// roleSvc := roleimpl.NewService(postgresDB, cfg)
	userSvc := userimpl.NewService(postgresDB, cfg)
	userAuthSvc := auth.NewService(userSvc, cfg)
	// movieSvc := movieimpl.NewService(postgresDB, cfg)
	// categorySvc := categoryimpl.NewService(postgresDB, cfg)
	// supplierSvc := supplierimpl.NewService(postgresDB, cfg)

	restServer := protocol.NewServer(&protocol.Dependencies{
		Postgres:       postgresDB,
		Cfg:            cfg,
		ContextHandler: contextHandler,

		UserSvc: userSvc,
		AuthSvc: userAuthSvc,
		// MovieSvc:         movieSvc,
		// RoleSvc:          roleSvc,
		AccessControlSvc: accesscontrolSvc,
		// CategorySvc:      categorySvc,
		// SupplierSvc:      supplierSvc,
	}, cfg)

	services := registry.NewServiceRegistry(
		restServer.Run,
	)

	if isStandaloneMode {
		services = registry.NewServiceRegistry(
			restServer.Run,
		)
	}

	return &Server{
		postgresDB: postgresDB,
		services:   services.GetServices(),
		log:        zap.L().Named("apiserver"),
	}, nil
}

func (s *Server) Run(ctx context.Context) {
	defer func() {
		s.postgresDB.Close()
	}()

	var wg sync.WaitGroup
	wg.Add(len(s.services))

	for _, svc := range s.services {
		go func(svc registry.RunFunc) error {
			defer wg.Done()
			err := svc(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				s.log.Error("stopped server", zap.String("service", serviceName), zap.Error(err))
				return fmt.Errorf("%s run error: %w", serviceName, err)
			}

			return nil
		}(svc)
	}

	wg.Wait()
}

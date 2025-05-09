package protocol

import (
	"backend/pkg/config"
	"backend/pkg/identity/accesscontrol"
	"backend/pkg/identity/auth"
	"backend/pkg/identity/user"
	"backend/pkg/infra/api/middleware"
	"backend/pkg/infra/api/routing"
	"backend/pkg/infra/storage/db"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	Dependencies    *Dependencies
	Router          *routing.Router
	log             *zap.Logger
	ShutdownTimeout time.Duration
}

type Dependencies struct {
	Postgres       db.DB
	Cfg            *config.Config
	ContextHandler *middleware.ContextHandler

	UserSvc user.Service
	AuthSvc auth.Service
	// MovieSvc         movie.Service
	// RoleSvc          role.Service
	AccessControlSvc accesscontrol.Service
	// CategorySvc      category.Service
	// SupplierSvc      supplier.Service
}

func NewServer(deps *Dependencies, cfg *config.Config) *Server {
	r := routing.NewRouter()
	return &Server{
		Dependencies:    deps,
		Router:          r,
		log:             zap.L().Named("server"),
		ShutdownTimeout: time.Duration(30) * time.Second,
	}
}

func (s *Server) registerRoutes() {
	r := s.Router

	s.NewUserHandler(r)
	s.NewAuthHandler(r)
	// s.NewMovieHandler(r)
	// s.NewRoleHandler(r)
	s.NewAccessControlHandler(r)
	// s.NewCategoryHandler(r)
	// s.NewSupplierHandler(r)
}

func (s *Server) addMiddleware() {
	r := s.Router

	r.Use(s.Dependencies.ContextHandler.Middleware)
}

func (s *Server) Run(ctx context.Context) error {
	stopCh := ctx.Done()

	s.addMiddleware()
	s.registerRoutes()

	addr := fmt.Sprintf(":%s", s.Dependencies.Cfg.Server.HTTPPort)

	s.log.Info("Starting HTTP server", zap.String("port", s.Dependencies.Cfg.Server.HTTPPort))
	err := s.Router.ListenToAddress(addr)
	if err != nil {
		return err
	}

	go func() {
		<-stopCh
		fmt.Println("Shutting down server...")
		s.Router.Shutdown(ctx)
	}()

	return nil
}

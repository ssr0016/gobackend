package authimpl

import (
	"backend/pkg/infra/api/auth"
	"backend/pkg/infra/api/auth/clients"
	"backend/pkg/infra/api/auth/jwt"
	"errors"

	"context"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	clients          map[string]auth.Client
	jwtTokenProvider jwt.TokenProvider
}

func New() (auth.Service, error) {
	jwtTokenProvider, err := jwt.NewTokenProvider()
	if err != nil {
		return nil, err
	}

	s := &service{
		clients:          make(map[string]auth.Client),
		jwtTokenProvider: jwtTokenProvider,
	}

	s.RegisterClient(clients.NewJWT(jwtTokenProvider))

	return s, nil
}

func (s *service) RegisterClient(c auth.Client) {
	s.clients[c.Name()] = c
}

func (s *service) JwtTokenProvider() jwt.TokenProvider {
	return s.jwtTokenProvider
}

func (s *service) Authenticate(ctx context.Context, r *fiber.Ctx) (*auth.SessionContext, error) {
	for _, item := range s.clients {
		session, err := s.authenticate(ctx, item, r)
		if err != nil {
			if errors.Is(err, fiber.ErrUnauthorized) {
				return nil, err
			}
			continue
		}

		if session != nil {
			return session, nil
		}
	}

	return &auth.SessionContext{}, nil
}

func (s *service) authenticate(ctx context.Context, c auth.Client, r *fiber.Ctx) (*auth.SessionContext, error) {
	session, err := c.Authenticate(ctx, r)
	if err != nil {
		return nil, err
	}

	return session, nil
}

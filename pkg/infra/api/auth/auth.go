package auth

import (
	"backend/pkg/infra/api/auth/jwt"
	"context"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Authenticate(ctx context.Context, r *fiber.Ctx) (*SessionContext, error)
	JwtTokenProvider() jwt.TokenProvider
}

type Client interface {
	Name() string
	Authenticate(ctx context.Context, r *fiber.Ctx) (*SessionContext, error)
}

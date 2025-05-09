package clients

import (
	"backend/pkg/infra/api/auth"
	"backend/pkg/infra/api/auth/jwt"
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type JWT struct {
	jwtProvider jwt.TokenProvider
}

func NewJWT(jwtProvider jwt.TokenProvider) *JWT {
	return &JWT{
		jwtProvider: jwtProvider,
	}
}

func (j *JWT) Name() string {
	return auth.ClientJWT
}

func (j *JWT) Authenticate(ctx context.Context, r *fiber.Ctx) (*auth.SessionContext, error) {
	header := r.Get("Authorization")
	if header == "" {
		fmt.Println("No Authorization header found")
		return nil, nil
	}

	parts := strings.Split(header, " ")
	if len(parts) < 2 || strings.ToLower(parts[0]) != "bearer" {
		fmt.Println("Invalid Authorization header format")
		return nil, nil
	}

	tokenPart := parts[1]
	if len(tokenPart) == 0 {
		fmt.Println("Invalid JWT token")
		return nil, fiber.ErrUnauthorized
	}

	claims, err := jwt.ValidateToken(tokenPart)
	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	if claims == nil {
		return nil, fiber.ErrUnauthorized
	}

	return &auth.SessionContext{
		UserID:   claims.UserID,
		Username: claims.Username,
		UserType: claims.UserType,
		// Permissions: claims.Permissions,
	}, nil
}

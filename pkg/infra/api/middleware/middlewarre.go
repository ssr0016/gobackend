package middleware

import (
	"backend/pkg/identity/accesscontrol"
	"backend/pkg/identity/account"
	"backend/pkg/infra/api/auth"
	"backend/pkg/infra/api/errors"
	"backend/pkg/infra/api/request"
	"backend/pkg/infra/api/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrUnauthorized = errors.New("authentication.unauthorized", "Unauthorized")
	ErrForbidden    = errors.New("authentication.forbidden", "Forbidden")
)

// ContextHandler manages authentication-related middleware
type ContextHandler struct {
	authSvc auth.Service
}

// New initializes and returns a ContextHandler with an auth service
func New(authSvc auth.Service) (*ContextHandler, error) {
	return &ContextHandler{authSvc: authSvc}, nil
}

// Middleware handles authentication and attaches the session to the Fiber request context
func (h *ContextHandler) Middleware(c *fiber.Ctx) error {
	reqCtx := &request.ReqContext{
		Session: &auth.SessionContext{},
	}

	ctx := request.WithReqContext(c.Context(), reqCtx)

	session, err := h.authSvc.Authenticate(ctx, c)
	if err != nil {
		fmt.Println("Authentication failed:", err) // Debug logging
		return response.SendError(c, fiber.StatusUnauthorized, ErrUnauthorized)
	}

	reqCtx.Session = session
	fmt.Println("Session is nil, blocking request")
	c.SetUserContext(ctx)

	return c.Next()
}

func ForUser(accountType account.Account) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("Middleware triggered for authentication")

		reqCtx, err := request.GetUserContext(c)
		if err != nil || reqCtx.Session == nil {
			fmt.Println("Unauthorized - No session found")
			return response.SendError(c, fiber.StatusUnauthorized, ErrUnauthorized)
		}

		if reqCtx.Session.UserType != accountType {
			fmt.Println("Forbidden - User type mismatch")
			return response.SendError(c, fiber.StatusForbidden, ErrForbidden)
		}

		fmt.Println("User authenticated:", reqCtx.Session.UserID)
		return c.Next()
	}
}

func ForPermission(requiredPermissions ...accesscontrol.Action) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// reqCtx, err := request.GetUserContext(c)
		// if err != nil || reqCtx.Session == nil {
		// 	return response.SendError(c, fiber.StatusUnauthorized, ErrUnauthorized)
		// }

		// for _, required := range requiredPermissions {
		// 	if !hasPermission(reqCtx.Session.Permissions, required) {
		// 		return response.SendError(c, fiber.StatusForbidden, ErrForbidden)
		// 	}
		// }

		return c.Next()
	}
}

func hasPermission(userPermissions []accesscontrol.Action, required accesscontrol.Action) bool {
	for _, perm := range userPermissions {
		if perm == required {
			return true
		}
	}
	return false
}

package request

import (
	"backend/pkg/identity/account"
	"backend/pkg/infra/api/auth"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ReqContext struct {
	Session *auth.SessionContext
}

type key int

const (
	userKey key = iota
	logKey
	requestKey
)

// Attach ReqContext to a parent context
func WithReqContext(parent context.Context, req *ReqContext) context.Context {
	return context.WithValue(parent, requestKey, req)
}

// Retrieve user session from context
func GetUserInfo(ctx context.Context) (auth.SessionContext, bool) {
	req, ok := ctx.Value(requestKey).(*ReqContext)
	if !ok || req == nil || req.Session == nil {
		return auth.SessionContext{}, false
	}
	return *req.Session, true
}

// Retrieve user session based on account type
func UserFromAccountType(ctx context.Context, accountType account.Account) (auth.SessionContext, bool) {
	req, ok := ctx.Value(requestKey).(*ReqContext)
	if !ok || req == nil || req.Session == nil {
		return auth.SessionContext{}, false
	}

	if req.Session.UserType != accountType { // Fix: Ensure we compare the correct field
		return auth.SessionContext{}, false
	}

	return *req.Session, true
}

// Get user context from Fiber's context
func GetUserContext(c *fiber.Ctx) (*ReqContext, error) {
	ctx := c.UserContext()
	reqCtx, ok := ctx.Value(requestKey).(*ReqContext)
	if !ok || reqCtx == nil {
		return nil, errors.New("user context not found or invalid")
	}
	return reqCtx, nil
}

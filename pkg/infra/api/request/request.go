package request

import (
	"backend/pkg/infra/api/auth"
	"context"
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

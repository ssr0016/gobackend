package auth

import "context"

type Service interface {
	Login(ctx context.Context, cmd *LoginCommand) (*LoginResult, error)
}

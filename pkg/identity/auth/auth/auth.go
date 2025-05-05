package auth

import (
	"backend/pkg/config"
	"backend/pkg/identity/account"
	"backend/pkg/identity/auth"
	"backend/pkg/identity/user"
	"backend/pkg/infra/api/auth/jwt"
	"backend/pkg/util/encrypt"
	"context"
	"time"

	"go.uber.org/zap"
)

type service struct {
	userSvc user.Service
	cfg     *config.Config
	log     *zap.Logger
}

func NewService(userSvc user.Service, cfg *config.Config) auth.Service {
	return &service{
		userSvc: userSvc,
		cfg:     cfg,
		log:     zap.L().Named("auth.service"),
	}
}

func (s *service) Login(ctx context.Context, cmd *auth.LoginCommand) (*auth.LoginResult, error) {
	exist, err := s.userSvc.GetByLoginName(ctx, cmd.LoginName)
	if err != nil {
		return nil, err
	}

	if exist == nil {
		return nil, auth.ErrUserOrPasswordInvalid
	}

	valid, err := encrypt.VerifyPassword(cmd.Password, exist.Salt, exist.Password)
	if err != nil {
		return nil, auth.ErrUserOrPasswordInvalid
	}

	if !valid {
		return nil, auth.ErrUserOrPasswordInvalid
	}

	authProvider, err := jwt.NewTokenProvider()
	if err != nil {
		return nil, err
	}

	accessToken, err := authProvider.GenerateToken(exist.ID, exist.LoginName, account.Admin, s.cfg.Auth.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	result := &auth.LoginResult{
		AuthToken: accessToken,
		LoginName: exist.LoginName,
		ExpiresAt: time.Now().UTC().Add(s.cfg.Auth.AccessTokenDuration).Unix(),
	}

	return result, nil
}

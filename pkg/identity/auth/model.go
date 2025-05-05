package auth

import (
	"backend/pkg/infra/api/errors"
)

var (
	ErrLoginNameRequired     = errors.New("auth.login-name-required", "Login name is required")
	ErrPasswordRequired      = errors.New("auth.password-required", "Password is required")
	ErrUserOrPasswordInvalid = errors.New("auth.user-or-password-invalid", "User or password is invalid")
)

type LoginCommand struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
}

type LoginResult struct {
	AuthToken string `json:"auth_token"`
	LoginName string `json:"login_name"`
	ExpiresAt int64  `json:"expires_at"`
}

func (c *LoginCommand) Validate() error {
	if len(c.LoginName) == 0 {
		return ErrLoginNameRequired
	}

	if len(c.Password) == 0 {
		return ErrPasswordRequired
	}

	return nil
}

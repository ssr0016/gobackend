package auth

import (
	"backend/pkg/identity/account"
)

type SessionContext struct {
	UserID   int64           `json:"user_id"`
	Username string          `json:"username"`
	UserType account.Account `json:"user_type"`
	// Permission []accesscontrol.Action `json:"permission"`
}

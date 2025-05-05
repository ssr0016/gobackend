package user

import "context"

type Service interface {
	Create(ctx context.Context, cmd *CreateUserCommand) error
	Search(ctx context.Context, query *SearchUserQuery) (*SearchUserResult, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, cmd *UpdateUserCommand) error
	UpdateStatus(ctx context.Context, cmd *UpdateStatusCommand) error
	UpdatePassword(ctx context.Context, cmd *UpdatePasswordCommand) error
	GetByLoginName(ctx context.Context, loginName string) (*User, error)
	// ForgotPassoword(ctx context.Context, cmd *ForgotPasswordCommand) error
	// AddUserToRole(ctx context.Context, cmd *AddUserToRoleCommand) error
	// GetPermissions(ctx context.Context, userID int64) ([]accesscontrol.Action, error)
}

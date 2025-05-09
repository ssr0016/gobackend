package accesscontrol

import "context"

type Service interface {
	GetPermissions(ctx context.Context) *AccessControlList
	GetPermissionsByRoleID(ctx context.Context, roleID int64) ([]Action, error)
}

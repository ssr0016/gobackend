package accesscontrolimpl

import (
	"backend/pkg/identity/accesscontrol"
	"backend/pkg/infra/storage/db"
	"context"

	"go.uber.org/zap"
)

type store struct {
	db  db.DB
	log *zap.Logger
}

func newStore(db db.DB) *store {
	return &store{
		db:  db,
		log: zap.L().Named("accesscontrol.store"),
	}
}

func (s *store) getPermissionsByRoleID(ctx context.Context, roleID int64) ([]accesscontrol.Action, error) {
	var result []accesscontrol.Action

	rawSQL := `
		SELECT
			action
		FROM "role_permission"
		WHERE
			role_id = ?
	`

	err := s.db.Select(ctx, &result, rawSQL, roleID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

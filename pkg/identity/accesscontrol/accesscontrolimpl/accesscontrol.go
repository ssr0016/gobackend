package accesscontrolimpl

import (
	"backend/pkg/config"
	"backend/pkg/identity/accesscontrol"
	"backend/pkg/infra/storage/db"
	"context"

	"go.uber.org/zap"
)

type service struct {
	store *store
	db    db.DB
	cfg   *config.Config
	log   *zap.Logger
}

func NewService(db db.DB, cfg *config.Config) accesscontrol.Service {
	return &service{
		store: newStore(db),
		db:    db,
		cfg:   cfg,
		log:   zap.L().Named("accesscontrol.service"),
	}
}

func (s *service) GetPermissions(ctx context.Context) *accesscontrol.AccessControlList {
	return &accesscontrol.AccessControls
}

func (s *service) GetPermissionsByRoleID(ctx context.Context, roleID int64) ([]accesscontrol.Action, error) {
	result, err := s.store.getPermissionsByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

package userimpl

import (
	"backend/pkg/config"
	"backend/pkg/identity/user"
	"backend/pkg/infra/storage/db"
	"backend/pkg/util/encrypt"
	"context"
	"time"

	"go.uber.org/zap"
)

type service struct {
	store *store
	db    db.DB
	cfg   *config.Config
	log   *zap.Logger
}

func NewService(
	db db.DB,
	cfg *config.Config,
) user.Service {
	return &service{
		store: newStore(db),
		db:    db,
		cfg:   cfg,
		log:   zap.L().Named("user.service"),
	}
}

func (s *service) Create(ctx context.Context, cmd *user.CreateUserCommand) error {
	now := time.Now().Format(time.RFC3339Nano)

	hashedPassword, err := encrypt.HashPassword(cmd.Password, cmd.Salt)
	if err != nil {
		return err
	}

	exist, err := s.store.isUserTaken(ctx, cmd.LoginName)
	if err != nil {
		return err
	}

	if exist {
		return user.ErrUserAlreadyExists
	}

	err = s.store.create(ctx, &user.User{
		UUID:       cmd.UUID,
		FirstName:  cmd.FirstName,
		LastName:   cmd.LastName,
		MiddleName: cmd.MiddleName,
		LoginName:  cmd.LoginName,
		Password:   hashedPassword,
		Status:     cmd.Status,
		Email:      cmd.Email,
		Salt:       cmd.Salt,
		CreatedBy:  "",
		CreatedAt:  now,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetByID(ctx context.Context, id int64) (*user.User, error) {
	result, err := s.store.getByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, user.ErrUserNotFound
	}

	return result, nil
}

func (s *service) Search(ctx context.Context, query *user.SearchUserQuery) (*user.SearchUserResult, error) {
	if query.Page <= 0 {
		query.Page = s.cfg.Pagination.Page
	}

	if query.PerPage <= 0 {
		query.PerPage = s.cfg.Pagination.PerPage
	}

	result, err := s.store.search(ctx, query)
	if err != nil {
		return nil, err
	}

	result.Page = query.Page
	result.PerPage = query.PerPage

	return result, nil
}

func (s *service) Update(ctx context.Context, cmd *user.UpdateUserCommand) error {
	now := time.Now().Format(time.RFC3339Nano)

	result, err := s.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = s.store.update(ctx, &user.User{
		ID:         result.ID,
		FirstName:  cmd.FirstName,
		LastName:   cmd.LastName,
		MiddleName: cmd.MiddleName,
		UpdatedAt:  &now,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateStatus(ctx context.Context, cmd *user.UpdateStatusCommand) error {
	now := time.Now().Format(time.RFC3339Nano)

	exist, err := s.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = s.store.updateStatus(ctx, &user.User{
		ID:        exist.ID,
		Status:    cmd.Status,
		UpdatedAt: &now,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdatePassword(ctx context.Context, cmd *user.UpdatePasswordCommand) error {
	now := time.Now().Format(time.RFC3339Nano)
	// signedInUser, _ := request.GetUserInfo(ctx)

	exist, err := s.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	hashedPassword, err := encrypt.HashPassword(cmd.Password, exist.Salt)
	if err != nil {
		return err
	}

	err = s.store.updatePassword(ctx, &user.User{
		ID:       exist.ID,
		Password: hashedPassword,
		// UpdatedBy: &signedInUser.Username,
		UpdatedAt: &now,
	})
	if err != nil {
		return err
	}

	return nil
}

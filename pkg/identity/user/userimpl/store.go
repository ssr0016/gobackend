package userimpl

import (
	"backend/pkg/identity/user"
	"backend/pkg/infra/storage/db"
	"bytes"
	"context"
	"database/sql"
	"errors"
	"strings"

	"go.uber.org/zap"
)

type store struct {
	db  db.DB
	log *zap.Logger
}

func newStore(db db.DB) *store {
	return &store{
		db:  db,
		log: zap.L().Named("user.store"),
	}
}

func (s *store) create(ctx context.Context, entity *user.User) error {
	rawSQL := `
			INSERT INTO "user" (
					uuid,
					first_name,
					last_name,
					middle_name,
					login_name,
					password,
					status,
					email,
					salt,
					created_by,
					created_at
			)
			VALUES (
				:uuid,
				:first_name,
				:last_name,
				:middle_name,
				:login_name,
				:password,
				:status,
				:email,
				:salt,
				:created_by,
				:created_at
			)
	`
	_, err := s.db.NamedExec(ctx, rawSQL, entity)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) isUserTaken(ctx context.Context, loginName string) (bool, error) {
	var result string

	rawSQL := `
		SELECT
			login_name
		FROM "user"
		WHERE
			login_name = ?
	`
	err := s.db.Get(ctx, &result, rawSQL, loginName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (s *store) getByID(ctx context.Context, id int64) (*user.User, error) {
	var result user.User

	rawSQL := `
		SELECT
			id,
			uuid,
			first_name,
			last_name,
			middle_name,
			login_name,
			password,
			status,
			email,
			salt,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM "user"
		WHERE
			id = ?
	`
	err := s.db.Get(ctx, &result, rawSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (s *store) search(ctx context.Context, query *user.SearchUserQuery) (*user.SearchUserResult, error) {
	var (
		result = &user.SearchUserResult{
			Users: make([]*user.User, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
	)

	sql.WriteString(`
		SELECT
			*
		FROM "user"
	`)

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE " + strings.Join(whereConditions, " AND "))
	}

	sql.WriteString(" ORDER BY created_at DESC")

	count, err := s.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(" LIMIT ? OFFSET ?")
		whereParams = append(whereParams, query.PerPage, offset)
	}

	err = s.db.Select(ctx, &result.Users, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil
}

func (s *store) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int64, error) {
	var count int64

	err := s.db.Get(ctx, &count, "SELECT COUNT(id) AS count FROM ("+sql.String()+") AS t1", whereParams...)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (s *store) update(ctx context.Context, entity *user.User) error {
	rawSQL := `
		UPDATE "user"
		SET
			first_name = :first_name,
			last_name = :last_name,
			middle_name = :middle_name,
			email = :email,
			updated_by = :updated_by,
			updated_at = :updated_at
		WHERE
			id = :id
	`
	_, err := s.db.NamedExec(ctx, rawSQL, entity)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) updateStatus(ctx context.Context, entity *user.User) error {
	rawSQL := `
		UPDATE "user"
		SET
			status = :status,
			updated_by = :updated_by,
			updated_at = :updated_at
		WHERE
			id = :id
	`
	_, err := s.db.NamedExec(ctx, rawSQL, entity)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) updatePassword(ctx context.Context, entity *user.User) error {
	rawSQL := `
		UPDATE "user"
		SET
			password = :password,
			updated_by = :updated_by,
			updated_at = :updated_at
		WHERE 
			id = :id
	`
	_, err := s.db.NamedExec(ctx, rawSQL, entity)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) getByLoginName(ctx context.Context, loginName string) (*user.User, error) {
	var result user.User

	rawSQL := `
		SELECT
			id,
			uuid,
			first_name,
			last_name,
			middle_name,
			login_name,
			password,
			status,
			email,
			salt,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM "user"
		WHERE
			login_name = ?
	`
	err := s.db.Get(ctx, &result, rawSQL, loginName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

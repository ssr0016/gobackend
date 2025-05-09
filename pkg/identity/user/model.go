package user

import (
	"backend/pkg/infra/api/errors"
	"backend/pkg/util/generator"
	"unicode"

	"github.com/google/uuid"
)

var (
	ErrInvalidLoginName      = errors.New("user.invalid-login-name", "Invalid login name")
	ErrInvalidPasswordLength = errors.New("user.invalid-password-length", "Invalid password length")
	ErrInvalidPassword       = errors.New("user.invalid-password", "Invalid password should contain at least one uppercase letter, number and special character")
	ErrInvalidFirstName      = errors.New("user.invalid-first-name", "Invalid first name")
	ErrInvalidLastName       = errors.New("user.invalid-last-name", "Invalid last name")
	ErrInvalidStatus         = errors.New("user.invalid-status", "Invalid status")
	ErrUserAlreadyExists     = errors.New("user.user-already-exists", "User already exists")
	ErrUserNotFound          = errors.New("user.user-not-found", "User not found")
	ErrInvalidUserID         = errors.New("user.invalid-user-id", "Invalid user id")
)

type Status int

const (
	Pending Status = iota + 1
	Active
	Inactive
	Deleted
)

type User struct {
	ID         int64   `db:"id" json:"id"`
	UUID       string  `db:"uuid" json:"uuid"`
	FirstName  string  `db:"first_name" json:"first_name"`
	LastName   string  `db:"last_name" json:"last_name"`
	MiddleName *string `db:"middle_name" json:"middle_name"`
	LoginName  string  `db:"login_name" json:"login_name"`
	Password   string  `db:"password" json:"-"`
	Status     Status  `db:"status" json:"status"`
	Email      *string `db:"email" json:"email"`
	Salt       string  `db:"salt" json:"-"`
	CreatedBy  string  `db:"created_by" json:"created_by"`
	CreatedAt  string  `db:"created_at" json:"created_at"`
	UpdatedBy  *string `db:"updated_by" json:"updated_by"`
	UpdatedAt  *string `db:"updated_at" json:"updated_at"`
}

type CreateUserCommand struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
	LoginName  string  `json:"login_name"`
	Email      *string `json:"email"`
	Password   string  `json:"password"`
	Status     Status  `json:"status"`
	UUID       string
	Salt       string
}

type SearchUserQuery struct {
	Page    int `schema:"page"`
	PerPage int `schema:"per_page"`
}

type SearchUserResult struct {
	TotalCount int64   `json:"total_count"`
	Users      []*User `json:"users"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

type UpdateUserCommand struct {
	ID         int64
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
}

type UpdateStatusCommand struct {
	ID     int64
	Status Status `json:"status"`
}

type UpdatePasswordCommand struct {
	ID       int64
	Password string `json:"password"`
	Salt     string
}

func (cmd *CreateUserCommand) Validate() error {
	if len(cmd.LoginName) == 0 {
		return ErrInvalidLoginName
	}

	err := ValidatePassword(cmd.Password)
	if err != nil {
		return err
	}

	if len(cmd.FirstName) == 0 {
		return ErrInvalidFirstName
	}

	if len(cmd.LastName) == 0 {
		return ErrInvalidLastName
	}

	if cmd.Status == 0 {
		cmd.Status = Pending
	}

	cmd.UUID = uuid.New().String()
	cmd.Salt, err = generator.GenerateSalt()
	if err != nil {
		return err
	}

	return nil
}

func (cmd *UpdateUserCommand) Validate() error {
	if cmd.ID <= 0 {
		return ErrInvalidUserID
	}

	if len(cmd.FirstName) == 0 {
		return ErrInvalidFirstName
	}

	if len(cmd.LastName) == 0 {
		return ErrInvalidLastName
	}

	return nil
}

func (cmd *UpdatePasswordCommand) Validate() error {
	if cmd.ID <= 0 {
		return ErrInvalidUserID
	}

	err := ValidatePassword(cmd.Password)
	if err != nil {
		return err
	}

	cmd.Salt, err = generator.GenerateSalt()
	if err != nil {
		return err
	}

	return nil
}

func (cmd *UpdateStatusCommand) Validate() error {
	if cmd.ID <= 0 {
		return ErrInvalidUserID
	}

	err := cmd.Status.Validate()
	if err != nil {
		return err
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrInvalidPasswordLength
	}

	var hasUppercase, hasSpecialChar, hasNumber bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	if !hasUppercase || !hasSpecialChar || !hasNumber {
		return ErrInvalidPassword
	}

	return nil
}

func (s Status) Validate() error {
	if s < Pending || s > Deleted {
		return ErrInvalidStatus
	}

	return nil
}

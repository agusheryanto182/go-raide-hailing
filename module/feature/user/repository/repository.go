package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/user"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/user/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
)

type userRepository struct {
	statements statements
}

// CheckConflict implements user.UserRepositoryInterface.
func (u *userRepository) CheckConflict(ctx context.Context, username string, email string, role string) error {
	var userExists, emailExists bool

	rows, err := u.statements.checkUsernameAndEmail.QueryxContext(ctx, username, email, role)
	if err != nil {
		return customErr.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&userExists, &emailExists); err != nil {
			return customErr.NewInternalServerError(err.Error())
		}
	}

	if userExists {
		return customErr.NewConflictError("username conflict with all types of users")
	}

	if emailExists {
		return customErr.NewConflictError("email conflict with another " + role)
	}

	return nil
}

// CheckUser implements user.UserRepositoryInterface.
func (u *userRepository) CheckUser(ctx context.Context, id string, username string, role string) (bool, error) {
	var userExists bool

	if err := u.statements.checkUser.QueryRowxContext(ctx, id, username, role).Scan(&userExists); err != nil {
		return false, customErr.NewInternalServerError(err.Error())
	}

	return userExists, nil
}

// FindByUsernameAndRole implements user.UserRepositoryInterface.
func (u *userRepository) FindByUsernameAndRole(ctx context.Context, username, role string) (*entities.User, error) {
	user := &entities.User{}

	rows, err := u.statements.findByUsernameAndRole.QueryxContext(ctx, username, role)
	if err != nil {
		return nil, customErr.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, customErr.NewNotFoundError("user not found")
			}
			return nil, customErr.NewInternalServerError(err.Error())
		}
	}

	return user, nil
}

// RegisterUser implements user.UserRepositoryInterface.
func (u *userRepository) RegisterUser(ctx context.Context, payload *dto.ReqCreateUser) (string, error) {
	var id string

	if err := u.statements.registerUser.QueryRowxContext(ctx, payload.Username, payload.Password, payload.Role, payload.Email).Scan(&id); err != nil {
		return "", customErr.NewInternalServerError(err.Error())
	}

	return id, nil
}

func NewUserRepository() user.UserRepositoryInterface {
	return &userRepository{
		statements: prepareStatements(),
	}
}

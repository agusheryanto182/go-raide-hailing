package repository

import (
	"context"

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

	if err := u.statements.checkUsernameAndEmail.QueryRowxContext(ctx, username, email, role).Scan(&userExists, &emailExists); err != nil {
		return customErr.NewInternalServerError(err.Error())
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

	if err := u.statements.findByUsernameAndRole.QueryRowxContext(ctx, username, role).Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, customErr.NewNotFoundError("user not found")
		}
		return nil, customErr.NewInternalServerError(err.Error())
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

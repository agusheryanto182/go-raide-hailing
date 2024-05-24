package service

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/feature/user"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/user/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/agusheryanto182/go-raide-hailing/utils/hash"
	"github.com/agusheryanto182/go-raide-hailing/utils/jwt"
)

type userService struct {
	userRepo user.UserRepositoryInterface
	jwt      jwt.JWTInterface
	hash     hash.HashInterface
}

// CheckUser implements user.UserServiceInterface.
func (u *userService) CheckUser(ctx context.Context, id string, username string, role string) (bool, error) {
	return u.userRepo.CheckUser(ctx, id, username, role)
}

// Login implements user.UserServiceInterface.
func (u *userService) Login(ctx context.Context, payload *dto.ReqLoginUser) (string, error) {
	user, err := u.userRepo.FindByUsernameAndRole(ctx, payload.Username, payload.Role)
	if err != nil {
		return "", err
	}

	if !u.hash.CheckPasswordHash(payload.Password, user.Password) {
		return "", customErr.NewBadRequestError("invalid credentials")
	}

	token, err := u.jwt.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		return "", customErr.NewInternalServerError(err.Error())
	}

	return token, nil
}

// RegisterUser implements user.UserServiceInterface.
func (u *userService) RegisterUser(ctx context.Context, payload *dto.ReqCreateUser) (string, error) {
	if err := u.userRepo.CheckConflict(ctx, payload.Username, payload.Email, payload.Role); err != nil {
		return "", err
	}

	hashed, err := u.hash.HashPassword(payload.Password)
	if err != nil {
		return "", customErr.NewInternalServerError(err.Error())
	}

	payload.Password = hashed

	id, err := u.userRepo.RegisterUser(ctx, payload)
	if err != nil {
		return "", err
	}

	token, err := u.jwt.GenerateJWT(id, payload.Username, payload.Role)
	if err != nil {
		return "", customErr.NewInternalServerError(err.Error())
	}

	return token, nil
}

func NewUserService(userRepo user.UserRepositoryInterface, jwt jwt.JWTInterface, hash hash.HashInterface) user.UserServiceInterface {
	return &userService{
		userRepo: userRepo,
		jwt:      jwt,
		hash:     hash,
	}
}

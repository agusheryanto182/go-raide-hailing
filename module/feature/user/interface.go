package user

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/user/dto"
	"github.com/gofiber/fiber/v2"
)

type UserRepositoryInterface interface {
	RegisterUser(ctx context.Context, payload *dto.ReqCreateUser) (string, error)
	FindByUsernameAndRole(ctx context.Context, username, role string) (*entities.User, error)
	CheckUser(ctx context.Context, id, username, role string) (bool, error)
	CheckConflict(ctx context.Context, username, email, role string) error
}

type UserServiceInterface interface {
	RegisterUser(ctx context.Context, payload *dto.ReqCreateUser) (string, error)
	Login(ctx context.Context, payload *dto.ReqLoginUser) (string, error)
	CheckUser(ctx context.Context, id, username, role string) (bool, error)
}

type UserControllerInterface interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

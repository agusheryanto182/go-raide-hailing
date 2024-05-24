package controller

import (
	"github.com/agusheryanto182/go-raide-hailing/module/feature/user"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/user/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/agusheryanto182/go-raide-hailing/utils/request"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService user.UserServiceInterface
}

// Login implements user.UserControllerInterface.
func (u *UserController) Login(ctx *fiber.Ctx) error {
	role := ctx.UserContext().Value("role").(string)
	req := new(dto.ReqLoginUser)
	req.Role = role

	if err := request.BindValidate(ctx, req); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	token, err := u.userService.Login(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// RegisterUser implements user.UserControllerInterface.
func (u *UserController) Register(ctx *fiber.Ctx) error {
	role := ctx.UserContext().Value("role").(string)
	req := new(dto.ReqCreateUser)

	if err := request.BindValidate(ctx, req); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	req.Role = role

	token, err := u.userService.RegisterUser(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func NewUserController(userService user.UserServiceInterface) user.UserControllerInterface {
	return &UserController{
		userService: userService,
	}
}

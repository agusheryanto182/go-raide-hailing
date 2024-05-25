package controller

import (
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/agusheryanto182/go-raide-hailing/utils/logging"
	"github.com/agusheryanto182/go-raide-hailing/utils/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type merchantController struct {
	merchantService merchant.MerchantServiceInterface
	validator       *validator.Validate
}

// CreateMerchant implements merchant.MerchantControllerInterface.
func (m *merchantController) CreateMerchant(ctx *fiber.Ctx) error {
	req := new(dto.ReqCreateMerchant)

	if err := request.BindValidate(ctx, req, m.validator); err != nil {
		logging.GetLogger("merchant").Debug(err.Error())
		return customErr.NewBadRequestError(err.Error())
	}

	id, err := m.merchantService.CreateMerchant(ctx.Context(), req)
	if err != nil {
		logging.GetLogger("merchant").Debug(err.Error())
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"merchantId": id,
	})
}

func NewMerchantController(merchantService merchant.MerchantServiceInterface, validator *validator.Validate) merchant.MerchantControllerInterface {
	return &merchantController{
		merchantService: merchantService,
		validator:       validator,
	}
}

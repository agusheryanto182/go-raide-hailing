package controller

import (
	"strings"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/agusheryanto182/go-raide-hailing/utils/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type merchantController struct {
	merchantService merchant.MerchantServiceInterface
	validator       *validator.Validate
}

// CreateMerchantItems implements merchant.MerchantControllerInterface.
func (m *merchantController) CreateMerchantItems(ctx *fiber.Ctx) error {
	merchantId := ctx.Params("merchantId")
	req := new(dto.ReqCreateMerchantItem)

	req.MerchantId = merchantId
	if err := request.BindValidate(ctx, req, m.validator); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	result, err := m.merchantService.CreateMerchantItems(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"itemId": result,
	})
}

// GetMerchantItemsByFilters implements merchant.MerchantControllerInterface.
func (m *merchantController) GetMerchantItemsByFilters(ctx *fiber.Ctx) error {
	merchantId := ctx.Params("merchantId")
	req := new(dto.ReqGetMerchantItemsByFilters)

	if err := ctx.QueryParser(req); err != nil {
		return customErr.NewBadRequestError("failed to parse query : " + err.Error())
	}

	req.MerchantId = merchantId

	cleanedItemId := strings.ReplaceAll(req.ItemId, "\"", "")
	req.ItemId = cleanedItemId

	if strings.ToUpper(req.CreatedAt) != "ASC" && strings.ToUpper(req.CreatedAt) != "DESC" {
		req.CreatedAt = "DESC"
	}

	if req.Limit == 0 {
		req.Limit = 5
	}

	meta, merchant_items, err := m.merchantService.GetMerchantItemsByFilters(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": merchant_items,
		"meta": fiber.Map{
			"limit":  meta.Limit,
			"offset": meta.Offset,
			"total":  meta.Total,
		},
	})
}

// GetMerchantByFilters implements merchant.MerchantControllerInterface.
func (m *merchantController) GetMerchantByFilters(ctx *fiber.Ctx) error {
	req := new(dto.ReqGetMerchantByFilters)

	if err := ctx.QueryParser(req); err != nil {
		return customErr.NewBadRequestError("failed to parse query : " + err.Error())
	}

	cleanedMerchantId := strings.ReplaceAll(req.MerchantId, "\"", "")
	req.MerchantId = cleanedMerchantId

	if req.Name != "" {
		cleanedName := strings.ReplaceAll(req.Name, "\"", "")
		req.Name = cleanedName
	}

	if req.MerchantCategory != "" {
		cleanedMerchantCategory := strings.ReplaceAll(req.MerchantCategory, "\"", "")
		if !entities.ValidCategoriesMerchant[cleanedMerchantCategory] {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"data": []*dto.ResGetMerchant{},
			})
		}
		req.MerchantCategory = cleanedMerchantCategory
	}

	if req.CreatedAt == "" {
		cleanedCreatedAt := strings.ReplaceAll(req.CreatedAt, "\"", "")
		if strings.ToUpper(cleanedCreatedAt) != "ASC" && strings.ToUpper(cleanedCreatedAt) != "DESC" {
			req.CreatedAt = "DESC"
		} else {
			req.CreatedAt = cleanedCreatedAt
		}
	}

	if req.Limit == 0 {
		req.Limit = 5
	}

	meta, merchants, err := m.merchantService.GetMerchantByFilters(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": merchants,
		"meta": fiber.Map{
			"limit":  meta.Limit,
			"offset": meta.Offset,
			"total":  meta.Total,
		},
	})
}

// CreateMerchant implements merchant.MerchantControllerInterface.
func (m *merchantController) CreateMerchant(ctx *fiber.Ctx) error {
	req := new(dto.ReqCreateMerchant)

	if err := request.BindValidate(ctx, req, m.validator); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	id, err := m.merchantService.CreateMerchant(ctx.Context(), req)
	if err != nil {
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

package controller

import (
	"strconv"
	"strings"

	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/agusheryanto182/go-raide-hailing/utils/jwt"
	"github.com/agusheryanto182/go-raide-hailing/utils/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type purchaseController struct {
	purchaseService purchase.PurchaseServiceInterface
	validator       *validator.Validate
}

// PostEstimate implements purchase.PurchaseControllerInterface.
func (p *purchaseController) PostEstimate(ctx *fiber.Ctx) error {
	currentUser := ctx.Locals("CurrentUser").(*jwt.JWTPayload)

	req := new(dto.ReqPostEstimate)
	req.UserId = currentUser.Id
	if err := request.BindValidate(ctx, req, p.validator); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	startingPoints := 0
	for _, order := range req.Orders {
		if order.IsStartingPoint {
			startingPoints++
		}
	}

	if startingPoints != 1 {
		return customErr.NewBadRequestError("Invalid starting point")
	}

	result, err := p.purchaseService.PostEstimate(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"totalPrice":                     result.TotalPrice,
		"estimatedDeliveryTimeInMinutes": result.EstimatedDeliveryTimeInMinutes,
		"calculatedEstimateId":           result.CalculatedEstimateId,
	})
}

// GetNearbyMerchants implements purchase.PurchaseControllerInterface.
func (p *purchaseController) GetNearbyMerchants(ctx *fiber.Ctx) error {
	coords := ctx.Params("coords")

	coordParts := strings.Split(coords, ",")
	if len(coordParts) != 2 {
		return customErr.NewBadRequestError("Invalid coordinates")
	}

	lat, err := strconv.ParseFloat(coordParts[0], 64)
	if err != nil {
		return customErr.NewBadRequestError("Invalid latitude : " + err.Error())
	}

	long, err := strconv.ParseFloat(coordParts[1], 64)
	if err != nil {
		return customErr.NewBadRequestError("Invalid longitude : " + err.Error())
	}

	req := new(dto.ReqNearbyMerchants)

	req.UserLat = lat
	req.UserLong = long

	if err := ctx.QueryParser(req); err != nil {
		return customErr.NewBadRequestError("failed to parse query : " + err.Error())
	}

	if req.MerchantCategory != "" {
		if err := p.validator.Struct(req); err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"data": []interface{}{},
			})
		}
	}

	cleanedMerchantId := strings.ReplaceAll(req.MerchantId, "\"", "")
	req.MerchantId = cleanedMerchantId

	if req.Limit == 0 {
		req.Limit = 5
	}

	results, err := p.purchaseService.GetNearbyMerchants(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": results,
	})
}

func NewPurchaseController(
	purchaseService purchase.PurchaseServiceInterface,
	validator *validator.Validate,
) purchase.PurchaseControllerInterface {
	return &purchaseController{
		purchaseService: purchaseService,
		validator:       validator,
	}
}

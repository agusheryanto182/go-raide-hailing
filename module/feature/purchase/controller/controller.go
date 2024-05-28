package controller

import (
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type purchaseController struct {
	validator *validator.Validate
}

// GetNearbyMerchants implements purchase.PurchaseControllerInterface.
func (p *purchaseController) GetNearbyMerchants(ctx *fiber.Ctx) error {
	// latStr := ctx.Params("lat")
	// longStr := ctx.Params("long")

	// lat, err := strconv.ParseFloat(latStr, 64)
	// if err != nil {
	// 	return customErr.NewBadRequestError("Invalid latitude")
	// }

	// long, err := strconv.ParseFloat(longStr, 64)
	// if err != nil {
	// 	return customErr.NewBadRequestError("Invalid longitude")
	// }
	panic("implement me")
}

func NewPurchaseController(validator *validator.Validate) purchase.PurchaseControllerInterface {
	return &purchaseController{
		validator: validator,
	}
}

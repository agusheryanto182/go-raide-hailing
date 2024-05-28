package routes

import (
	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/image"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/user"
	"github.com/agusheryanto182/go-raide-hailing/module/middleware"
	"github.com/agusheryanto182/go-raide-hailing/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, controller user.UserControllerInterface, jwtService jwt.JWTInterface, userService user.UserServiceInterface) {
	admin := app.Group("/admin")
	admin.Post("/register", middleware.AddRoleToContext(entities.RoleAdmin), controller.Register)
	admin.Post("/login", middleware.AddRoleToContext(entities.RoleAdmin), controller.Login)

	user := app.Group("/users")
	user.Post("/register", middleware.AddRoleToContext(entities.RoleUser), controller.Register)
	user.Post("/login", middleware.AddRoleToContext(entities.RoleUser), controller.Login)
}

func ImageRoute(app *fiber.App, controller image.ImageControllerInterface, jwtService jwt.JWTInterface, userService user.UserServiceInterface) {
	image := app.Group("/image")
	image.Post("", middleware.ProtectedWithRole(jwtService, userService, entities.RoleAdmin), controller.UploadImage)
}

func MerchantRoute(app *fiber.App, controller merchant.MerchantControllerInterface, controllerPurchase purchase.PurchaseControllerInterface, jwtService jwt.JWTInterface, userService user.UserServiceInterface) {
	merchantAdmin := app.Group("/admin/merchants")
	merchantAdmin.Post("", middleware.ProtectedWithRole(jwtService, userService, entities.RoleAdmin), controller.CreateMerchant)
	merchantAdmin.Get("", middleware.ProtectedWithRole(jwtService, userService, entities.RoleAdmin), controller.GetMerchantByFilters)
	merchantAdmin.Post("/:merchantId/items", middleware.ProtectedWithRole(jwtService, userService, entities.RoleAdmin), controller.CreateMerchantItems)
	merchantAdmin.Get("/:merchantId/items", middleware.ProtectedWithRole(jwtService, userService, entities.RoleAdmin), controller.GetMerchantItemsByFilters)

	merchantUser := app.Group("/merchants")
	merchantUser.Get("/nearby/:coords", middleware.ProtectedWithRole(jwtService, userService, entities.RoleUser), controllerPurchase.GetNearbyMerchants)
}

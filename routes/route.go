package routes

import (
	"github.com/agusheryanto182/go-raide-hailing/module/feature/image"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/user"
	"github.com/agusheryanto182/go-raide-hailing/module/middleware"
	"github.com/agusheryanto182/go-raide-hailing/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, controller user.UserControllerInterface, jwtService jwt.JWTInterface, userService user.UserServiceInterface) {
	admin := app.Group("/admin")
	admin.Post("/register", middleware.AddRoleToContext("admin"), controller.Register)
	admin.Post("/login", middleware.AddRoleToContext("admin"), controller.Login)

	user := app.Group("/users")
	user.Post("/register", middleware.AddRoleToContext("user"), controller.Register)
	user.Post("/login", middleware.AddRoleToContext("user"), controller.Login)
}

func ImageRoute(app *fiber.App, controller image.ImageControllerInterface, jwtService jwt.JWTInterface, userService user.UserServiceInterface) {
	image := app.Group("/image")
	image.Post("", middleware.Protected(jwtService, userService), controller.UploadImage)
}

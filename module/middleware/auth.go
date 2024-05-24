package middleware

import (
	"context"
	"strings"

	"github.com/agusheryanto182/go-raide-hailing/module/feature/user"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/agusheryanto182/go-raide-hailing/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected(jwtService jwt.JWTInterface, userService user.UserServiceInterface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			return customErr.NewUnauthorizedError("Access denied: invalid token")
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		payload, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return customErr.NewUnauthorizedError("Access denied: invalid token")
		}

		user, err := userService.CheckUser(context.Background(), payload.Id, payload.Username, payload.Role)
		if err != nil {
			return err
		}

		if !user {
			return customErr.NewUnauthorizedError("Access denied: invalid token")
		}

		c.Locals("CurrentUser", payload)

		return c.Next()
	}
}

func AddRoleToContext(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.WithValue(c.UserContext(), "role", role)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"ww-api-gateway/api/presenter"
	"ww-api-gateway/pkg/util"
)

func Protected(validationKey string) fiber.Handler {
	vk, err := util.GetJwtPublicKey(validationKey)
	if err != nil {
		panic(err)
	}
	return jwtware.New(jwtware.Config{
		SigningKey:    vk,
		SigningMethod: "RS256",
		ErrorHandler:  jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrorResponse(err))
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrorResponse(err))
}

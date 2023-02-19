package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/rs/zerolog/log"
	"ww-api/api/presenter"
	"ww-api/pkg/util"
)

func Protected(validationKey string) fiber.Handler {
	vk, err := util.GetJwtPublicKey(validationKey)
	if err != nil {
		log.Err(err).Msg("failed to get jwt public key")
		return nil
	}
	return jwtware.New(jwtware.Config{
		SigningKey:    vk,
		SigningMethod: "RS256",
		ErrorHandler:  jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		log.Err(err).Msg("missing or malformed jwt")
		return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrorResponse(err))
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrorResponse(err))
}

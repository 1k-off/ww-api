package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"ww-api/api/presenter"
	"ww-api/pkg/entities"
	"ww-api/pkg/target"
)

func GetCheckerTargets(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		switch name {
		case entities.CheckerNameUptime:
			targets, err := svc.GetTargetsForUptimeChecker()
			if err != nil {
				log.Err(err).Msg("failed to get targets for uptime checker")
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.CheckerErrorResponse(name, err))
			}
			return c.JSON(presenter.CheckerUptimeSuccessResponse(targets))
		case entities.CheckerNameSsl:
			targets, err := svc.GetTargetsForSslChecker()
			if err != nil {
				log.Err(err).Msg("failed to get targets for ssl checker")
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.CheckerErrorResponse(name, err))
			}
			return c.JSON(presenter.CheckerSslSuccessResponse(targets))
		case entities.CheckerNameDomainExpiration:
			targets, err := svc.GetTargetsForDomainExpirationChecker()
			if err != nil {
				log.Err(err).Msg("failed to get targets for domain expiration checker")
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.CheckerErrorResponse(name, err))
			}
			return c.JSON(presenter.CheckerDomainExpirationSuccessResponse(targets))
		default:
			log.Err(errors.New("invalid checker name")).Msgf("invalid checker name %s", name)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CheckerErrorResponse(name, errors.New("invalid checker name")))
		}
	}
}

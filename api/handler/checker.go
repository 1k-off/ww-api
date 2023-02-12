package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
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
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.CheckerErrorResponse(name, err))
			}
			return c.JSON(presenter.CheckerUptimeSuccessResponse(targets))
		case entities.CheckerNameSsl:
			targets, err := svc.GetTargetsForSslChecker()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.CheckerErrorResponse(name, err))
			}
			return c.JSON(presenter.CheckerSslSuccessResponse(targets))
		case entities.CheckerNameDomainExpiration:
			targets, err := svc.GetTargetsForDomainExpirationChecker()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.CheckerErrorResponse(name, err))
			}
			return c.JSON(presenter.CheckerDomainExpirationSuccessResponse(targets))
		default:
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CheckerErrorResponse(name, errors.New("invalid checker name")))
		}
	}
}

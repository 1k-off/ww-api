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
		var checkerName string
		switch name {
		case entities.CheckerNameUptime:
			checkerName = entities.CheckerNameUptime
		case entities.CheckerNameSsl:
			checkerName = entities.CheckerNameSsl
		case entities.CheckerNameDomainExpiration:
			checkerName = entities.CheckerNameDomainExpiration
		default:
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CheckerErrorResponse(name, errors.New("invalid checker name")))
		}
		t, err := svc.GetTargetsForChecker(checkerName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.CheckerErrorResponse(name, err))
		}
		return c.JSON(presenter.CheckerSuccessResponse(name, t))
	}
}

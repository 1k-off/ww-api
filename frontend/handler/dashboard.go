package handler

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/pkg/auth"
)

func Dashboard(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := checkAccessToken(c, svc); err == nil {
			return c.Render("dashboard", fiber.Map{})
		}
		return c.Redirect("/", fiber.StatusSeeOther)
	}
}

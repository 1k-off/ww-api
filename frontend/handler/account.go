package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"ww-api/pkg/auth"
	"ww-api/pkg/user"
)

func AccountGet(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := checkAccessToken(c, svc); err == nil {
			user, err := svc.GetCurrentUser(c.Cookies("access_token"))
			if err != nil {
				return err
			}
			return c.Render("account", fiber.Map{
				"Username": user.Login,
			})
		}
		return c.Redirect("/", fiber.StatusSeeOther)
	}
}

func AccountPost(authSvc auth.Service, userSvc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := checkAccessToken(c, authSvc); err == nil {
			user, err := authSvc.GetCurrentUser(c.Cookies("access_token"))
			if err != nil {
				log.Err(err).Msg("failed to get current user")
				return err
			}
			newPassword := c.FormValue("new-password")
			if newPassword != "" {
				user.Password = newPassword
				user, err = userSvc.Update(user)
				if err != nil {
					log.Err(err).Msg("failed to update user")
					return err
				}
				log.Debug().Msgf("user %s updated", user.Login)
				return c.Render("account", fiber.Map{
					"Username": user.Login,
				})
			}
		}
		return c.Redirect("/", fiber.StatusSeeOther)
	}
}

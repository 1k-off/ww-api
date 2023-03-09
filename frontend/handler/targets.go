package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"strconv"
	"ww-api/pkg/auth"
	"ww-api/pkg/target"
)

func TargetsGet(authSvc auth.Service, targetSvc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := checkAccessToken(c, authSvc); err == nil {
			targets, err := targetSvc.GetAll()
			if err != nil {
				return c.Render("error", fiber.Map{
					"Error": err.Error(),
				})
			}
			return c.Render("targets", fiber.Map{
				"Targets": targets,
			})
		}
		return c.Redirect("/", fiber.StatusSeeOther)
	}
}

func TargetsPost(authSvc auth.Service, targetSvc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := checkAccessToken(c, authSvc); err == nil {
			target, err := targetSvc.GetByUrl(c.FormValue("url"))
			if err != nil {
				return err
			}
			isActive, err := parseBool(c.FormValue("is-active"))
			if err != nil {
				log.Err(err).Msgf("error parsing is-active: %v", c.FormValue("is-active"))
				return err
			}
			uptime, err := parseBool(c.FormValue("uptime"))
			if err != nil {
				log.Err(err).Msgf("error parsing uptime: %v", c.FormValue("uptime"))
				return err
			}
			ssl, err := parseBool(c.FormValue("ssl"))
			if err != nil {
				log.Err(err).Msgf("error parsing ssl: %v", c.FormValue("ssl"))
				return err
			}
			domainExpiration, err := parseBool(c.FormValue("domain-expiration"))
			if err != nil {
				log.Err(err).Msgf("error parsing domain-expiration: %v", c.FormValue("domain-expiration"))
				return err
			}
			target.IsActive = isActive
			target.Uptime = uptime
			target.SSL = ssl
			target.DomainExpiration = domainExpiration
			log.Debug().Msgf("target: %v", target)
			_, err = targetSvc.Update(target)
			if err != nil {
				return err
			}
			targets, _ := targetSvc.GetAll()
			return c.Render("targets", fiber.Map{
				"Targets": targets,
			})
		}
		return c.Redirect("/", fiber.StatusSeeOther)
	}
}

func parseBool(value string) (bool, error) {
	if value == "on" {
		return true, nil
	}
	if value == "off" || value == "" {
		return false, nil
	}
	return strconv.ParseBool(value)
}

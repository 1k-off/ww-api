package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"ww-api/frontend/presenter"
	"ww-api/pkg/auth"
	"ww-api/pkg/entities"
)

const (
	cookieNameAccessToken  = "access_token"
	cookieNameRefreshToken = "refresh_token"
)

func LoginGet(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := checkAccessToken(c, svc); err == nil {
			return c.Redirect("/dashboard", fiber.StatusFound)
		}
		return c.Render("login", fiber.Map{})
	}
}

func LoginPost(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// parse form data
		login := c.FormValue("username")
		password := c.FormValue("password")
		if login == "" || password == "" {
			log.Debug().Msgf("failed to parse form data %s", login)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrorResponse(errors.New("invalid data")))
		}
		u, accessToken, refreshToken, err := svc.Login(&entities.User{
			Login:    login,
			Password: password,
		})
		if err != nil {
			log.Err(err).Msgf("failed to login user %s", u.Login)
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrorResponse(err))
		}
		accessTokenCookie := &fiber.Cookie{
			Name:    cookieNameAccessToken,
			Value:   accessToken,
			Expires: svc.AccessTokenExpiresIn(),
		}
		refreshTokenCookie := &fiber.Cookie{
			Name:    cookieNameRefreshToken,
			Value:   refreshToken,
			Expires: svc.RefreshTokenExpiresIn(),
		}
		c.Cookie(accessTokenCookie)
		c.Cookie(refreshTokenCookie)
		log.Info().Msgf("user %s logged in", u.Login)
		return c.Redirect("/dashboard", fiber.StatusFound)
	}
}

func checkAccessToken(c *fiber.Ctx, svc auth.Service) error {
	if c.Cookies(cookieNameAccessToken) != "" {
		ok, err := svc.CheckAccessToken(c.Cookies(cookieNameAccessToken))
		if err != nil {
			log.Err(err).Msg("failed to check access token")
			return err
		}
		if ok {
			return nil
		}
	}
	return errors.New("invalid access token")
}

// Logout deletes the access and refresh tokens from the cookies
func Logout(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.ClearCookie(cookieNameAccessToken)
		c.ClearCookie(cookieNameRefreshToken)
		return c.Redirect("/", fiber.StatusSeeOther)
	}
}

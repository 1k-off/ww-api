package handler

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/api/presenter"
	"ww-api/pkg/auth"
	"ww-api/pkg/entities"
)

const (
	cookieNameAccessToken   = "access_token"
	cookieNameRefreshToken  = "refresh_token"
	headerNameAuthorization = "Authorization"
)

func Login(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := new(entities.User)
		if err := c.BodyParser(u); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrorResponse(err))
		}
		u, accessToken, refreshToken, err := svc.Login(u)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.AuthErrorResponse(err))
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
		return c.JSON(presenter.AuthSuccessResponse(u, accessToken, refreshToken))
	}
}

func Refresh(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refreshToken := c.Get(cookieNameRefreshToken)
		u, accessToken, refreshToken, err := svc.Refresh(refreshToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.AuthErrorResponse(err))
		}
		return c.JSON(presenter.AuthSuccessResponse(u, accessToken, refreshToken))
	}
}

func Logout(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// getting authorization header content and splitting it to get the token
		authHeader := c.Get(headerNameAuthorization)
		accessToken := authHeader[7:]
		err := svc.Logout(accessToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.AuthErrorResponse(err))
		}
		c.ClearCookie(cookieNameAccessToken, cookieNameRefreshToken)
		return c.JSON(presenter.AuthLogoutResponse())
	}
}

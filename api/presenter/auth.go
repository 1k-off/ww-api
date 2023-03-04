package presenter

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/pkg/entities"
)

type AuthSuccess struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func AuthSuccessResponse(u *entities.User, accessToken, refreshToken string) *AuthSuccess {
	return &AuthSuccess{
		ID:           u.ID.Hex(),
		Login:        u.Login,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func AuthErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"error":  err.Error(),
	}
}

func AuthLogoutResponse() *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   nil,
	}
}

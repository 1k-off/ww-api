package presenter

import (
	"github.com/gofiber/fiber/v2"
	"ww-api-gateway/pkg/entities"
)

type AuthSuccess struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func AuthSuccessResponse(u *entities.User, accessToken, refreshToken string) *fiber.Map {
	auth := &AuthSuccess{
		ID:           u.ID.Hex(),
		Login:        u.Login,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   auth,
	}
}

func AuthErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"error":  err.Error(),
		"data":   nil,
	}
}

func AuthLogoutResponse() *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   nil,
	}
}

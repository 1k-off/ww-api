package presenter

import (
	"github.com/gofiber/fiber/v2"
	"ww-api-gateway/pkg/entities"
)

type UserResponse struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

func UserSuccessResponse(u *entities.User) *fiber.Map {
	user := &UserResponse{
		ID:    u.ID.Hex(),
		Login: u.Login,
	}
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   user,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"error":  err.Error(),
		"data":   nil,
	}
}

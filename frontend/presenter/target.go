package presenter

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/pkg/entities"
)

func TargetSuccessResponse(t *entities.Target) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   t,
	}
}

func TargetErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"error":  err.Error(),
		"data":   nil,
	}
}

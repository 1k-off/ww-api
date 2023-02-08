package presenter

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/pkg/entities"
)

func MetricsDownSuccessResponse(d []*entities.TargetDown) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   d,
	}
}
func MetricsSslExpirationSuccessResponse(d []*entities.SslExpiringSoon) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   d,
	}
}
func MetricsDomainExpirationSuccessResponse(d []*entities.DomainExpiringSoon) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   d,
	}
}

func MetricsStatsSuccessResponse(d *entities.MetricsStats) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   d,
	}
}

func MetricsPutSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   nil,
	}
}

func MetricsErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"error":  err.Error(),
		"data":   nil,
	}
}

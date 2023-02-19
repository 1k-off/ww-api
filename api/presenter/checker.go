package presenter

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/pkg/entities"
)

func CheckerUptimeSuccessResponse(targets []*entities.UptimeTarget) []*entities.UptimeTarget {
	return targets
}

func CheckerSslSuccessResponse(targets []*entities.SslTarget) []*entities.SslTarget {
	return targets
}
func CheckerDomainExpirationSuccessResponse(targets []*entities.DomainExpirationTarget) []*entities.DomainExpirationTarget {
	return targets
}

func CheckerErrorResponse(name string, err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"error":  err.Error(),
		"name":   name,
		"data":   nil,
	}
}

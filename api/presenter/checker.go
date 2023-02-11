package presenter

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/pkg/entities"
)

type UptimeSuccessResponse struct {
	ID     string                 `json:"id"`
	URL    string                 `json:"url"`
	Config *entities.UptimeConfig `json:"config"`
}

type SslSuccessResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type DomainExpirationSuccessResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func CheckerSuccessResponse(name string, targets []*entities.Target) interface{} {
	switch name {
	case entities.CheckerNameUptime:
		return CheckerUptimeSuccessResponse(targets)
	case entities.CheckerNameSsl:
		return CheckerSslSuccessResponse(targets)
	case entities.CheckerNameDomainExpiration:
		return CheckerDomainExpirationSuccessResponse(targets)
	}
	return nil
}

func CheckerUptimeSuccessResponse(targets []*entities.Target) *fiber.Map {
	var uptimeTargets []*UptimeSuccessResponse
	for _, target := range targets {
		uptimeTargets = append(uptimeTargets, &UptimeSuccessResponse{
			ID:     target.ID.Hex(),
			URL:    target.URL,
			Config: &target.Config.Uptime,
		})
	}
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   uptimeTargets,
	}
}
func CheckerSslSuccessResponse(targets []*entities.Target) []*SslSuccessResponse {
	var sslTargets []*SslSuccessResponse
	for _, target := range targets {
		sslTargets = append(sslTargets, &SslSuccessResponse{
			ID:  target.ID.Hex(),
			URL: target.URL,
		})
	}
	return sslTargets
}
func CheckerDomainExpirationSuccessResponse(targets []*entities.Target) *fiber.Map {
	var domainExpirationTargets []*DomainExpirationSuccessResponse
	for _, target := range targets {
		domainExpirationTargets = append(domainExpirationTargets, &DomainExpirationSuccessResponse{
			ID:  target.ID.Hex(),
			URL: target.URL,
		})
	}
	return &fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   domainExpirationTargets,
	}
}

func CheckerErrorResponse(name string, err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"error":  err.Error(),
		"name":   name,
		"data":   nil,
	}
}

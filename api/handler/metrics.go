package handler

import (
	"github.com/gofiber/fiber/v2"
	"ww-api-gateway/api/presenter"
	"ww-api-gateway/pkg/entities"
	"ww-api-gateway/pkg/metrics"
)

func MetricsGetDown(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d, err := svc.GetDownTargets()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsDownSuccessResponse(d))
	}
}

func MetricsGetSslExpirationSoon(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d, err := svc.GetSslExpiringSoon()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsSslExpirationSuccessResponse(d))
	}
}

func MetricsGetDomainExpirationSoon(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d, err := svc.GetDomainExpiringSoon()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsDomainExpirationSuccessResponse(d))
	}
}

func MetricsGetStats(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d, err := svc.GetStats()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsStatsSuccessResponse(d))
	}
}

func MetricsPutUptime(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var m []*entities.UptimeData
		if err := c.BodyParser(&m); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.MetricsErrorResponse(err))
		}
		if err := svc.InsertUptime(m); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsPutSuccessResponse())
	}
}

func MetricsPutSslExpiration(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var m []*entities.SslData
		if err := c.BodyParser(&m); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.MetricsErrorResponse(err))
		}
		if err := svc.InsertSsl(m); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsPutSuccessResponse())
	}
}

func MetricsPutDomainExpiration(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var m []*entities.DomainExpirationData
		if err := c.BodyParser(&m); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.MetricsErrorResponse(err))
		}
		if err := svc.InsertDomainExpiration(m); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsPutSuccessResponse())
	}
}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"ww-api/api/presenter"
	"ww-api/pkg/entities"
	"ww-api/pkg/metrics"
)

func MetricsGetDown(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d, err := svc.GetDownTargets()
		if err != nil {
			log.Err(err).Msg("failed to get down targets")
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsDownSuccessResponse(d))
	}
}

func MetricsGetSslExpirationSoon(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d, err := svc.GetSslExpiringSoon()
		if err != nil {
			log.Err(err).Msg("failed to get ssl expiring soon")
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsSslExpirationSuccessResponse(d))
	}
}

func MetricsGetDomainExpirationSoon(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d, err := svc.GetDomainExpiringSoon()
		if err != nil {
			log.Err(err).Msg("failed to get domain expiring soon")
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsDomainExpirationSuccessResponse(d))
	}
}

func MetricsGetStats(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d, err := svc.GetStats()
		if err != nil {
			log.Err(err).Msg("failed to get stats")
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsStatsSuccessResponse(d))
	}
}

func MetricsPutUptime(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var m []*entities.UptimeData
		if err := c.BodyParser(&m); err != nil {
			log.Err(err).Msg("failed to parse body for uptime metrics")
			return c.Status(fiber.StatusBadRequest).JSON(presenter.MetricsErrorResponse(err))
		}
		if err := svc.InsertUptimeBatch(m); err != nil {
			log.Err(err).Msg("failed to batch insert uptime metrics")
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsPutSuccessResponse())
	}
}

func MetricsPutSslExpiration(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var m []*entities.SslData
		if err := c.BodyParser(&m); err != nil {
			log.Err(err).Msg("failed to parse body for ssl metrics")
			return c.Status(fiber.StatusBadRequest).JSON(presenter.MetricsErrorResponse(err))
		}
		if err := svc.InsertSslBatch(m); err != nil {
			log.Err(err).Msg("failed to batch insert ssl metrics")
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsPutSuccessResponse())
	}
}

func MetricsPutDomainExpiration(svc metrics.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var m []*entities.DomainExpirationData
		if err := c.BodyParser(&m); err != nil {
			log.Err(err).Msg("failed to parse body for domain expiration metrics")
			return c.Status(fiber.StatusBadRequest).JSON(presenter.MetricsErrorResponse(err))
		}
		if err := svc.InsertDomainExpirationBatch(m); err != nil {
			log.Err(err).Msg("failed to batch insert domain expiration metrics")
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.MetricsErrorResponse(err))
		}
		return c.JSON(presenter.MetricsPutSuccessResponse())
	}
}

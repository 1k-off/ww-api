package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api-gateway/api/handler"
	"ww-api-gateway/api/middleware"
	"ww-api-gateway/pkg/metrics"
)

func MetricsRouter(r fiber.Router, svc metrics.Service, validationKey string) {
	r.Get("/metrics/down", middleware.Protected(validationKey), handler.MetricsGetDown(svc))
	r.Get("/metrics/ssl-expiration-soon", middleware.Protected(validationKey), handler.MetricsGetSslExpirationSoon(svc))
	r.Get("/metrics/domain-expiration-soon", middleware.Protected(validationKey), handler.MetricsGetDomainExpirationSoon(svc))
	r.Get("/metrics/stats", middleware.Protected(validationKey), handler.MetricsGetStats(svc))
	r.Put("/metrics/uptime", middleware.Protected(validationKey), handler.MetricsPutUptime(svc))
	r.Put("/metrics/ssl", middleware.Protected(validationKey), handler.MetricsPutSslExpiration(svc))
	r.Put("/metrics/domain-expiration", middleware.Protected(validationKey), handler.MetricsPutDomainExpiration(svc))
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api-gateway/api/handler"
	"ww-api-gateway/api/middleware"
	"ww-api-gateway/pkg/target"
)

func TargetRouter(r fiber.Router, svc target.Service, validationKey string) {
	r.Get("/target/:id", middleware.Protected(validationKey), handler.TargetGet(svc))
	r.Post("/target", middleware.Protected(validationKey), handler.TargetAdd(svc))
	r.Delete("/target/:id", middleware.Protected(validationKey), handler.TargetDelete(svc))
	r.Patch("/target/:id", middleware.Protected(validationKey), handler.TargetUpdate(svc))
	r.Get("/targets", middleware.Protected(validationKey), handler.TargetsGetAll(svc))
	//r.Get("/targets/down", middleware.Protected(validationKey), handler.TargetsGetDown(svc))
	//r.Get("/targets/ssl", middleware.Protected(validationKey), handler.TargetsGetSslExp(svc))
	//r.Get("/targets/domain-expiration", middleware.Protected(validationKey), handler.TargetsGetDomainExp(svc))
	//r.Get("/targets/stats", middleware.Protected(validationKey), handler.TargetsGetStats(svc))
}

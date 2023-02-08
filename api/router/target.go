package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/api/handler"
	"ww-api/api/middleware"
	"ww-api/pkg/target"
)

func TargetRouter(r fiber.Router, svc target.Service, validationKey string) {
	r.Get("/target/:id", middleware.Protected(validationKey), handler.TargetGet(svc))
	r.Post("/target", middleware.Protected(validationKey), handler.TargetAdd(svc))
	r.Delete("/target/:id", middleware.Protected(validationKey), handler.TargetDelete(svc))
	r.Patch("/target/:id", middleware.Protected(validationKey), handler.TargetUpdate(svc))
	r.Get("/targets", middleware.Protected(validationKey), handler.TargetsGetAll(svc))
}

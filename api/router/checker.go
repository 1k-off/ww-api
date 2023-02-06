package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api-gateway/api/handler"
	"ww-api-gateway/api/middleware"
	"ww-api-gateway/pkg/target"
)

func CheckerRouter(r fiber.Router, svc target.Service, validationKey string) {
	r.Get("/checker/:name", middleware.Protected(validationKey), handler.GetCheckerTargets(svc))
}

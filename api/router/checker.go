package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/api/handler"
	"ww-api/api/middleware"
	"ww-api/pkg/target"
)

func CheckerRouter(r fiber.Router, svc target.Service, validationKey string) {
	r.Get("/checker/:name", middleware.Protected(validationKey), handler.GetCheckerTargets(svc))
}

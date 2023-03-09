package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/frontend/handler"
	"ww-api/pkg/auth"
)

func Dashboard(r fiber.Router, svc auth.Service) {
	r.Get("/dashboard", handler.Dashboard(svc))
}

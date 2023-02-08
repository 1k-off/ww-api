package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/api/handler"
	"ww-api/api/middleware"
	"ww-api/pkg/auth"
)

func AuthRouter(r fiber.Router, svc auth.Service, validationKey string) {
	r.Post("/auth/login", handler.Login(svc))
	r.Post("/auth/refresh", handler.Refresh(svc))
	r.Post("/auth/logout", middleware.Protected(validationKey), handler.Logout(svc))
}

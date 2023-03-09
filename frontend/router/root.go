package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/frontend/handler"
	"ww-api/pkg/auth"
)

func Root(r fiber.Router, svc auth.Service) {
	r.Get("/", handler.LoginGet(svc))
	r.Post("/", handler.LoginPost(svc))
	r.Get("/logout", handler.Logout(svc))
}

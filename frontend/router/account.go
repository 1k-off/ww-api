package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/frontend/handler"
	"ww-api/pkg/auth"
	"ww-api/pkg/user"
)

func Account(r fiber.Router, authSvc auth.Service, userSvc user.Service) {
	r.Get("/account", handler.AccountGet(authSvc))
	r.Post("/account", handler.AccountPost(authSvc, userSvc))
}

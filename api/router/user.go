package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api-gateway/api/handler"
	"ww-api-gateway/pkg/user"
)

func UserRouter(r fiber.Router, svc user.Service) {
	r.Get("/user/:id", handler.GetUser(svc))
	r.Post("/user", handler.CreateUser(svc))
	r.Put("/user/:id", handler.UpdateUser(svc))
	r.Delete("/user/:id", handler.DeleteUser(svc))
}

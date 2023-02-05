package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api-gateway/api/handler"
	"ww-api-gateway/api/middleware"
	"ww-api-gateway/pkg/user"
)

func UserRouter(r fiber.Router, svc user.Service, validationKey string) {
	r.Get("/user/:id", middleware.Protected(validationKey), handler.GetUser(svc))
	r.Post("/user", middleware.Protected(validationKey), handler.CreateUser(svc))
	r.Put("/user/:id", middleware.Protected(validationKey), handler.UpdateUser(svc))
	r.Delete("/user/:id", middleware.Protected(validationKey), handler.DeleteUser(svc))
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"ww-api/frontend/handler"
	"ww-api/pkg/auth"
	"ww-api/pkg/target"
)

func Targets(r fiber.Router, authSvc auth.Service, targetSvc target.Service) {
	r.Get("/targets", handler.TargetsGet(authSvc, targetSvc))
	r.Post("/targets", handler.TargetsPost(authSvc, targetSvc))
}

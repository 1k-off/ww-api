package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog/log"
	"ww-api/api/router"
	"ww-api/pkg/app"
)

func Start(s *app.Service, port string) error {
	server := fiber.New()
	server.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${ua} - ${status} - ${method} ${path}\n",
	}))
	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Wazzup, man!")
	})

	api := server.Group("/api")
	router.UserRouter(api, s.UserService, s.AccessTokenPublicKey)
	router.AuthRouter(api, s.AuthService, s.AccessTokenPublicKey)
	router.TargetRouter(api, s.TargetService, s.AccessTokenPublicKey)
	router.CheckerRouter(api, s.TargetService, s.AccessTokenPublicKey)
	router.MetricsRouter(api, s.MetricsService, s.AccessTokenPublicKey)
	log.Info().Msg("api server started")
	return server.Listen(":" + port)
}

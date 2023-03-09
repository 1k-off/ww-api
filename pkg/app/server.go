package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
	apiRouter "ww-api/api/router"
	"ww-api/frontend"
	frontendRouter "ww-api/frontend/router"
)

func StartServer(s *Service, port string) error {
	engine := frontend.NewEngine()
	server := fiber.New(fiber.Config{
		Views: engine,
	})
	server.Static("/assets/css", "./frontend/assets/css")
	server.Static("/assets/webfonts", "./frontend/assets/webfonts")
	server.Static("/assets/img", "./frontend/assets/img")

	server.Use(recover.New())
	server.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${ua} - ${status} - ${method} ${path}\n",
	}))
	//server.Use(cors.New(cors.Config{
	//	AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
	//	AllowOrigins:     "*",
	//	AllowCredentials: true,
	//	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	//}))
	//server.Get("/", func(ctx *fiber.Ctx) error {
	//	return ctx.SendString("Wazzup, man!")
	//})

	fn := server.Group("/")
	frontendRouter.Root(fn, s.AuthService)
	frontendRouter.Dashboard(fn, s.AuthService)
	frontendRouter.Targets(fn, s.AuthService, s.TargetService)
	frontendRouter.Account(fn, s.AuthService, s.UserService)

	api := server.Group("/api")
	apiRouter.UserRouter(api, s.UserService, s.AccessTokenPublicKey)
	apiRouter.AuthRouter(api, s.AuthService, s.AccessTokenPublicKey)
	apiRouter.TargetRouter(api, s.TargetService, s.AccessTokenPublicKey)
	apiRouter.CheckerRouter(api, s.TargetService, s.AccessTokenPublicKey)
	apiRouter.MetricsRouter(api, s.MetricsService, s.AccessTokenPublicKey)
	log.Info().Msg("server started")
	return server.Listen(":" + port)
}

package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewServer() *fiber.App {
	appCfg := fiber.Config{
		ReadTimeout:              15 * time.Second,
		EnableSplittingOnParsers: true,
		Immutable:                true,
	}

	app := fiber.New(appCfg)

	setMiddlewares(app)

	setRoutes(app)

	return app
}

func setMiddlewares(app *fiber.App) {
	useCorsMiddleware(app)
	useLogger(app)
}

func useLogger(app *fiber.App) {
	app.Use(logger.New())
}

func useCorsMiddleware(app *fiber.App) {
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: "*", //TODO: ajust this when we have some infra
		// AllowCredentials: true,
		AllowMethods:  "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:  "Content-Type, Correlation-Id",
		ExposeHeaders: "Content-Type, Correlation-Id",
	})

	app.Use(corsMiddleware)
}

func setRoutes(app *fiber.App) {
	router := newRouter(app)

	router.SetMinutaRoutes()
}

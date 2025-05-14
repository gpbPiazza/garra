package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gpbPiazza/garra/infra/envs"
	"github.com/gpbPiazza/garra/infra/http/health"
	"github.com/gpbPiazza/garra/infra/http/minuta"
)

func NewServer() *fiber.App {
	appCfg := fiber.Config{
		ReadTimeout:              15 * time.Second,
		EnableSplittingOnParsers: true,
		Immutable:                true,
	}

	app := fiber.New(appCfg)

	health.SetRoutes(app)

	setMiddlewares(app)

	setRoutes(app)

	return app
}

func setMiddlewares(app *fiber.App) {
	useCorsMiddleware(app)
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
}

func useCorsMiddleware(app *fiber.App) {
	configCors := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length,Content-Type",
	}

	if envs.IsProduction() {
		envsVar := envs.GetEnvs()
		configCors.AllowOrigins = envsVar.AllowOriginsHost
		configCors.AllowCredentials = true
	}

	corsMiddleware := cors.New(configCors)
	app.Use(corsMiddleware)
}

func setRoutes(app *fiber.App) {
	router := newRouter(app)
	minuta.SetRoutes(router.apiV1)
}

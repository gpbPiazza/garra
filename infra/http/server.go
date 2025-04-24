package http

import (
	"fmt"
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
	d := newDispatcher(app)

	setMinutaRoutes(d)
}

const (
	internalPrefix = "internal"
	clientPrefix   = "api"
)

type Dispatcher struct {
	// IternalRoute is dedicated to internal APIs  calls inside of our infra network.
	InternalV1Router fiber.Router

	// ClientV1Router is the router dedicated to any HTTP request from public-internet.
	// And must be autenticated.
	ClientV1Router fiber.Router

	// ClientV1Router is the router dedicated to any HTTP request from public-internet.
	// Not autentication is needed
	// Not implemented yeat
	PublicV1Router fiber.Router
}

func newDispatcher(app *fiber.App) Dispatcher {
	internalAPI := app.Group(fmt.Sprintf("/%s/v1", internalPrefix))
	clientAPI := app.Group(fmt.Sprintf("/%s/v1", clientPrefix))

	return Dispatcher{
		InternalV1Router: internalAPI,
		ClientV1Router:   clientAPI,
	}
}

package http

import "github.com/gofiber/fiber/v2"

type router struct {
	app        *fiber.App
	apiV1      fiber.Router
	internalV1 fiber.Router
}

func newRouter(app *fiber.App) *router {
	return &router{
		app:        app,
		apiV1:      app.Group("/api/v1"),
		internalV1: app.Group("/internal/v1"),
	}
}

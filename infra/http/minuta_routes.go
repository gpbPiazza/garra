package http

import "github.com/gofiber/fiber/v2"

const minutaRoutePrefix = "/minuta"

func setMinutaRoutes(d Dispatcher) {
	d.ClientV1Router.Route(minutaRoutePrefix, internalMinutaRoutes(), "client_routes")
}

func internalMinutaRoutes() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("", PostMinutaHandler)
	}
}

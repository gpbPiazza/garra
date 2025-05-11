package minuta

import "github.com/gofiber/fiber/v2"

func SetRoutes(apiV1 fiber.Router) {
	apiV1.Post("/generator/minuta", PostGeneratorHandler)
}

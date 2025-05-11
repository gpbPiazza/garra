package health

import "github.com/gofiber/fiber/v2"

func SetRoutes(apiV1 fiber.Router) {
	apiV1.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("I'm Alive")
	})
}

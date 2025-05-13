package health

import "github.com/gofiber/fiber/v2"

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("I'm Alive")
	})
}

package http

import (
	"errors"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/alemao-bigodes/application/minuta"
)

func PostMinutaHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("pdf")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to upload file")
	}

	if file.Header.Get("Content-Type") != "application/pdf" {
		return c.Status(fiber.StatusBadRequest).SendString("Only PDF files are allowed")
	}

	app := minuta.NewGeneratorApp()

	f, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("server error")
	}

	var allDoc string
	for !errors.Is(err, io.EOF) {
		buff := make([]byte, 1024) // 1 KB

		_, err := f.Read(buff)
		if err != nil {
			log.Printf("err while reading file err: %s", err)
		}

		allDoc += string(buff)
	}

	result, err := app.Generate(allDoc)
	if err != nil {
		log.Printf("err to generate minuta err: %s", err)
		return c.Status(fiber.StatusInternalServerError).SendString("server error")
	}

	c.Response().Header.Add("Content-Type", "text/html")

	return c.Status(fiber.StatusOK).SendString(result)
}

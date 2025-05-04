package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/alemao-bigodes/application/minuta"
	"github.com/gpbPiazza/alemao-bigodes/infra/pdf"
)

func PostGeneratorMinutaHandler(c *fiber.Ctx) error {
	formFile, err := c.FormFile("ato_consultar_pdf")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to upload file")
	}

	if formFile.Header.Get("Content-Type") != "application/pdf" {
		return c.Status(fiber.StatusBadRequest).SendString("Only PDF files are allowed")
	}

	minutaGenerator := minuta.NewGeneratorApp()

	pdfContentStr := pdf.ContentStr(formFile)

	result, err := minutaGenerator.Generate(pdfContentStr)
	if err != nil {
		log.Printf("err to generate minuta err: %s", err)
		return c.Status(fiber.StatusInternalServerError).SendString("server error")
	}

	c.Response().Header.Add("Content-Type", "text/html")
	return c.Status(fiber.StatusOK).SendString(result)
}

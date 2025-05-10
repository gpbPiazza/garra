package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/garra/application/minuta"
	"github.com/gpbPiazza/garra/infra/pdf"
)

type MinutaGeneratorBody struct {
	IsTransmitenteOverqualified bool `json:"is_transmitente_overqualified"`
	IsAdquirenteOverqualified   bool `json:"is_adquirente_overqualified"`
}

func (b MinutaGeneratorBody) PDFKey() string {
	return "ato_consultar_pdf"
}

func PostGeneratorMinutaHandler(c *fiber.Ctx) error {
	var body MinutaGeneratorBody

	formFile, err := c.FormFile(body.PDFKey())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to upload file")
	}

	if formFile.Header.Get("Content-Type") != "application/pdf" {
		return c.Status(fiber.StatusBadRequest).SendString("Only PDF files are allowed")
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("error to parse body - %s", err))
	}

	minutaGenerator := minuta.NewGeneratorApp()

	pdfContentStr := pdf.ContentStr(formFile)

	params := minuta.GenerateParams{
		DocStr:                      pdfContentStr,
		IsTransmitenteOverqualified: body.IsTransmitenteOverqualified,
		IsAdquirenteOverqualified:   body.IsAdquirenteOverqualified,
	}

	result, err := minutaGenerator.Generate(params)
	if err != nil {
		log.Printf("err to generate minuta err: %s", err)
		return c.Status(fiber.StatusInternalServerError).SendString("server error")
	}

	c.Response().Header.Add("Content-Type", "text/html")
	return c.Status(fiber.StatusOK).SendString(result)
}

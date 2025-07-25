package minuta

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/garra/application/minuta"
	"github.com/gpbPiazza/garra/infra/http/response"
	"github.com/gpbPiazza/garra/infra/pdf"
)

const (
	errCodeDOC = "#malformed_doc"
)

func PostGeneratorHandler(c *fiber.Ctx) error {
	formFile, err := c.FormFile("ato_consultar_pdf")
	if err != nil {
		return response.BadRequest(c, response.ErrorInfo{
			Code:    errCodeDOC,
			Message: "Failed to find form file with key ato_consultar_pdf"},
		)
	}

	if formFile.Header.Get("Content-Type") != "application/pdf" {
		return response.BadRequest(c, response.ErrorInfo{
			Code:    errCodeDOC,
			Message: "Missing expected Content-Type header"},
		)
	}

	is_transmitente_overqualified := c.FormValue("is_transmitente_overqualified", "false")
	is_adquirente_overqualified := c.FormValue("is_adquirente_overqualified", "false")

	minutaGenerator := minuta.NewGeneratorApp()

	pdfContentStr := pdf.ContentStr(formFile)

	params := minuta.GenerateParams{
		DocStr:               pdfContentStr,
		IsTransOverqualified: is_transmitente_overqualified == "true",
		IsAdquiOverqualified: is_adquirente_overqualified == "true",
	}

	resp, err := minutaGenerator.Generate(params)
	if err != nil {
		return response.InternalError(c)
	}

	return response.OK(c, resp)
}

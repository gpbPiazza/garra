package http

import (
	"log"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/alemao-bigodes/application/minuta"
	"github.com/ledongthuc/pdf"
)

func PostMinutaHandler(c *fiber.Ctx) error {
	formFile, err := c.FormFile("pdf")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to upload file")
	}

	if formFile.Header.Get("Content-Type") != "application/pdf" {
		return c.Status(fiber.StatusBadRequest).SendString("Only PDF files are allowed")
	}

	app := minuta.NewGeneratorApp()

	allDoc := parseDocToStr(formFile)

	result, err := app.Generate(allDoc)
	if err != nil {
		log.Printf("err to generate minuta err: %s", err)
		return c.Status(fiber.StatusInternalServerError).SendString("server error")
	}

	c.Response().Header.Add("Content-Type", "text/html")

	return c.Status(fiber.StatusOK).SendString(result)
}

func parseDocToStr(formFile *multipart.FileHeader) string {
	file, err := formFile.Open()
	if err != nil {
		log.Fatalf("err to open PDF err: %s", err)
	}

	fileReader, err := pdf.NewReader(file, formFile.Size)
	if err != nil {
		log.Fatalf("err to create new Redaer PDF err: %s", err)
	}

	log.Printf("file name: %s", formFile.Filename)
	log.Printf("file size: %d bytes", formFile.Size)

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("err to close file err: %s", err)
		}
	}()

	var allDoc string
	for pIndex := 1; pIndex <= fileReader.NumPage(); pIndex++ {
		page := fileReader.Page(pIndex)
		if page.V.IsNull() {
			log.Printf("page %d - isNull", pIndex)
		}

		pText, err := page.GetPlainText(nil)
		if err != nil {
			log.Fatalf("err at page %d - on GetPlainText err: %s", pIndex, err)
		}
		allDoc += "\n" + pText
	}

	return allDoc
}

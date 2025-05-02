package pdf

import (
	"log"
	"mime/multipart"

	"github.com/ledongthuc/pdf"
)

// ContentStr will receive a formFile and get all content in the file and return as string.
// ContentStr will breakline between pages in the pdf file.
func ContentStr(formFile *multipart.FileHeader) string {
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

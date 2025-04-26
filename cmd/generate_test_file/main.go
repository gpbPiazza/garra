package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/ledongthuc/pdf"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		log.Printf("args len %d", len(args))
		log.Printf("args: %v", args)
		log.Fatal("you must provide exactly 2 arguments: 1 - sourceFilepath, 2 - destFileName")
	}

	filePath := args[1]
	fileDestName := args[2]

	file, fReader, err := pdf.Open(filePath)
	if err != nil {
		log.Fatalf("error to open file err: %s", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("err to close file err: %s", err)
		}
	}()

	var allDoc string
	for pIndex := 1; pIndex <= fReader.NumPage(); pIndex++ {
		page := fReader.Page(pIndex)
		if page.V.IsNull() {
			log.Printf("page %d - isNull", pIndex)
		}

		pText, err := page.GetPlainText(nil)
		if err != nil {
			log.Fatalf("err at page %d - on GetPlainText err: %s", pIndex, err)
		}
		allDoc += "\n" + pText
	}

	err = os.WriteFile(fmt.Sprintf("./infra/test_files/%s.txt", fileDestName), []byte(allDoc), fs.ModePerm)
	if err != nil {
		log.Fatalf("err to write test file err: %s", err)
	}

	log.Print("finish with sucess!")
}

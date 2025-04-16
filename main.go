package main

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/ledongthuc/pdf"
)

func main() {
	file, r, err := pdf.Open("./assets/ato_consultar_ato.pdf")
	if err != nil {
		log.Fatalf("err to open PDF err: %s", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("err to close file err: %s", err)
		}
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("err to se file info err: %s", err)
	}

	logFileInfo(fileInfo)

	for pIndex := 1; pIndex <= r.NumPage(); pIndex++ {
		page := r.Page(1)
		if page.V.IsNull() {
			continue
		}

		result, err := page.GetPlainText(nil)
		if err != nil {
			log.Fatalf("err on GetPlainText err: %s", err)
		}
		fmt.Println(result)
	}

	fmt.Println("Finish reading file")
}

// func isSameSentence(text pdf.Text, lastTextStyle pdf.Text) bool {
// return (text.Font == lastTextStyle.Font) && (text.FontSize == lastTextStyle.FontSize)
//&& strings.Contains(lastTextStyle.S, ".")
// }

func logFileInfo(info fs.FileInfo) {
	log.Printf("file name: %s", info.Name())
	log.Printf("file mode: %s", info.Mode())
	log.Printf("file size: %d bytes", info.Size())
}

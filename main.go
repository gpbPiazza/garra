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

	// extractor := minuta.NewExtractor()

	for pIndex := 1; pIndex <= r.NumPage(); pIndex++ {
		page := r.Page(pIndex)
		if page.V.IsNull() {
			log.Printf("page %d - isNull", pIndex)
		}

		pText, err := page.GetPlainText(nil)
		if err != nil {
			log.Fatalf("err at page %d - on GetPlainText err: %s", pIndex, err)
		}
		fmt.Println(pText)
		// extractor.Extract(pText)
	}

	// valuesByReplaceKeys := extractor.Result()

	// params := minuta.MinutaParams{
	// 	Transmitente:          valuesByReplaceKeys[minuta.Transmitente],
	// 	Adquirente:            valuesByReplaceKeys[minuta.Adquirente],
	// 	TitleAto:              valuesByReplaceKeys[minuta.TitleAto],
	// 	TabelionatoNum:        valuesByReplaceKeys[minuta.TabelionatoNum],
	// 	TabelionatoName:       valuesByReplaceKeys[minuta.TabelionatoName],
	// 	TabelionatoCityState:  valuesByReplaceKeys[minuta.TabelionatoCityState],
	// 	BookNum:               valuesByReplaceKeys[minuta.BookNum],
	// 	BookPages:             valuesByReplaceKeys[minuta.BookPages],
	// 	EscrituraMadeDate:     valuesByReplaceKeys[minuta.EscrituraMadeDate],
	// 	EscrituraValor:        valuesByReplaceKeys[minuta.EscrituraValor],
	// 	EscrituraValorExtenso: valuesByReplaceKeys[minuta.EscrituraValorExtenso],
	// 	ItbiValor:             valuesByReplaceKeys[minuta.ItbiValor],
	// 	ItbiIncidenciaValor:   valuesByReplaceKeys[minuta.ItbiIncidenciaValor],
	// }

	// generatedMin := minuta.Minuta(params)

	fmt.Println("Minuta gerada")
	// fmt.Println(generatedMin)
	fmt.Println("Finish Minuta generator file")
}

func logFileInfo(info fs.FileInfo) {
	log.Printf("file name: %s", info.Name())
	log.Printf("file mode: %s", info.Mode())
	log.Printf("file size: %d bytes", info.Size())
}

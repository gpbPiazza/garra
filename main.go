package main

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/gpbPiazza/alemao-bigodes/extractor"
	"github.com/gpbPiazza/alemao-bigodes/minuta"
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

	ex := extractor.New()

	var allDoc string
	for pIndex := 1; pIndex <= r.NumPage(); pIndex++ {
		page := r.Page(pIndex)
		if page.V.IsNull() {
			log.Printf("page %d - isNull", pIndex)
		}

		pText, err := page.GetPlainText(nil)
		if err != nil {
			log.Fatalf("err at page %d - on GetPlainText err: %s", pIndex, err)
		}
		allDoc += "\n" + pText
	}

	ex.Extract(allDoc)
	dataExtracted := ex.Result()

	params := minuta.MinutaParams{
		TitleAto:             dataExtracted[extractor.TitleAto],
		TabelionatoName:      dataExtracted[extractor.TabelionatoName],
		TabelionatoCityState: dataExtracted[extractor.TabelionatoCityState],
		BookNum:              dataExtracted[extractor.BookNum],
		EscrituraMadeDate:    dataExtracted[extractor.EscrituraMadeDate],
		EscrituraValor:       dataExtracted[extractor.EscrituraValor],
		ItbiValor:            dataExtracted[extractor.ItbiValor],
		ItbiIncidenciaValor:  dataExtracted[extractor.ItbiIncidenciaValor],
		Transmitente: minuta.PersonParams{
			Name:          dataExtracted[extractor.OutorganteName],
			Job:           dataExtracted[extractor.OutorganteJob],
			Nationality:   dataExtracted[extractor.OutorganteNationality],
			MaritalStatus: dataExtracted[extractor.OutorganteEstadoCivil],
			CPF_CNPJ:      dataExtracted[extractor.OutorganteCPF_CNPJ],
			Address: minuta.AddressParams{
				Rua:          dataExtracted[extractor.OutorganteEnderecoRua],
				Num:          dataExtracted[extractor.OutorganteEnderecoN],
				CityUF:       dataExtracted[extractor.OutorganteEnderecoCidadeUF],
				Neighborhood: dataExtracted[extractor.OutorganteEnderecoBairro],
			},
		},
		Adquirente: minuta.PersonParams{
			Name:          dataExtracted[extractor.OutorgadoName],
			Job:           dataExtracted[extractor.OutorgadoJob],
			Nationality:   dataExtracted[extractor.OutorgadoNationality],
			MaritalStatus: dataExtracted[extractor.OutorgadoEstadoCivil],
			CPF_CNPJ:      dataExtracted[extractor.OutorgadoCPF_CNPJ],
			Address: minuta.AddressParams{
				Rua:          dataExtracted[extractor.OutorgadoEnderecoRua],
				Num:          dataExtracted[extractor.OutorgadoEnderecoN],
				CityUF:       dataExtracted[extractor.OutorgadoEnderecoCidadeUF],
				Neighborhood: dataExtracted[extractor.OutorgadoEnderecoBairro],
			},
		},
		InitialBookPages: dataExtracted[extractor.InitialBookPages],
		FinalBookPages:   dataExtracted[extractor.FinalBookPages],
	}

	generatedMin := minuta.Minuta(params)

	fmt.Println("Minuta gerada")
	fmt.Println(generatedMin)
	fmt.Println("Finish Minuta generator file")
}

func logFileInfo(info fs.FileInfo) {
	log.Printf("file name: %s", info.Name())
	log.Printf("file mode: %s", info.Mode())
	log.Printf("file size: %d bytes", info.Size())
}

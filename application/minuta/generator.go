package minuta

import (
	"github.com/gpbPiazza/alemao-bigodes/domain/extractor"
	"github.com/gpbPiazza/alemao-bigodes/domain/minuta"
)

type GeneratorApp struct {
}

func NewGeneratorApp() *GeneratorApp {
	return &GeneratorApp{}
}

func (app *GeneratorApp) Generate(docStr string) (string, error) {
	ex := extractor.New()

	ex.Extract(docStr)

	dataExtracted := ex.Result()

	params := minuta.MinutaParams{
		TitleAto:            dataExtracted[extractor.TitleAto],
		TabelionatoName:     dataExtracted[extractor.TabelionatoName],
		TabelionatoCityUF:   dataExtracted[extractor.TabelionatoCityUF],
		BookNum:             dataExtracted[extractor.BookNum],
		EscrituraMadeDate:   dataExtracted[extractor.EscrituraMadeDate],
		EscrituraValor:      dataExtracted[extractor.EscrituraValor],
		ItbiValor:           dataExtracted[extractor.ItbiValor],
		ItbiIncidenciaValor: dataExtracted[extractor.ItbiIncidenciaValor],
		Transmitente: minuta.PersonParams{
			Name:            dataExtracted[extractor.OutorganteName],
			Job:             dataExtracted[extractor.OutorganteJob],
			Nationality:     dataExtracted[extractor.OutorganteNationality],
			MaritalStatus:   dataExtracted[extractor.OutorganteEstadoCivil],
			DocNum_CPF_CNPJ: dataExtracted[extractor.OutorganteDocNumCPF_CNPJ],
			DocType:         dataExtracted[extractor.OutorganteDocType],
			Address: minuta.AddressParams{
				Rua:          dataExtracted[extractor.OutorganteEnderecoRua],
				Num:          dataExtracted[extractor.OutorganteEnderecoN],
				CityUF:       dataExtracted[extractor.OutorganteEnderecoCidadeUF],
				Neighborhood: dataExtracted[extractor.OutorganteEnderecoBairro],
			},
		},
		Adquirente: minuta.PersonParams{
			Name:            dataExtracted[extractor.OutorgadoName],
			Job:             dataExtracted[extractor.OutorgadoJob],
			Nationality:     dataExtracted[extractor.OutorgadoNationality],
			MaritalStatus:   dataExtracted[extractor.OutorgadoEstadoCivil],
			DocNum_CPF_CNPJ: dataExtracted[extractor.OutorgadoDocNumCPF_CNPJ],
			DocType:         dataExtracted[extractor.OutorgadoDocType],
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

	return minuta.Minuta(params), nil
}

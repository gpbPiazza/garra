package minuta

import (
	"github.com/gpbPiazza/garra/domain/extractor"
	"github.com/gpbPiazza/garra/domain/minuta"
)

type GeneratorApp struct {
}

func NewGeneratorApp() *GeneratorApp {
	return &GeneratorApp{}
}

type GenerateParams struct {
	DocStr                      string
	IsTransmitenteOverqualified bool
	IsAdquirenteOverqualified   bool
}

func (app *GeneratorApp) Generate(generateParams GenerateParams) (string, error) {
	ex := extractor.New()

	ex.Extract(generateParams.DocStr)

	extracted := ex.Result()

	params := minuta.MinutaParams{
		TitleAto:            extracted.Result[extractor.TitleAto],
		TabelionatoName:     extracted.Result[extractor.TabelionatoName],
		TabelionatoCityUF:   extracted.Result[extractor.TabelionatoCityUF],
		BookNum:             extracted.Result[extractor.BookNum],
		EscrituraMadeDate:   extracted.Result[extractor.EscrituraMadeDate],
		EscrituraValor:      extracted.Result[extractor.EscrituraValor],
		ItbiValor:           extracted.Result[extractor.ItbiValor],
		ItbiIncidenciaValor: extracted.Result[extractor.ItbiIncidenciaValor],
		Transmitente: minuta.PersonParams{
			IsOverqualified: generateParams.IsTransmitenteOverqualified,
			Name:            extracted.Result[extractor.OutorganteName],
			Nationality:     extracted.Result[extractor.OutorganteNationality],
			MaritalStatus:   extracted.Result[extractor.OutorganteEstadoCivil],
			DocNum_CPF_CNPJ: extracted.Result[extractor.OutorganteDocNumCPF_CNPJ],
			DocType:         extracted.Result[extractor.OutorganteDocType],
			Sex:             extracted.Result[extractor.OutorganteSex],
			Address: minuta.AddressParams{
				Rua:          extracted.Result[extractor.OutorganteEnderecoRua],
				Num:          extracted.Result[extractor.OutorganteEnderecoN],
				CityUF:       extracted.Result[extractor.OutorganteEnderecoCidadeUF],
				Neighborhood: extracted.Result[extractor.OutorganteEnderecoBairro],
			},
		},
		Adquirente: minuta.PersonParams{
			IsOverqualified: generateParams.IsAdquirenteOverqualified,
			Name:            extracted.Result[extractor.OutorgadoName],
			Nationality:     extracted.Result[extractor.OutorgadoNationality],
			MaritalStatus:   extracted.Result[extractor.OutorgadoEstadoCivil],
			DocNum_CPF_CNPJ: extracted.Result[extractor.OutorgadoDocNumCPF_CNPJ],
			DocType:         extracted.Result[extractor.OutorgadoDocType],
			Sex:             extracted.Result[extractor.OutorgadoSex],
			Address: minuta.AddressParams{
				Rua:          extracted.Result[extractor.OutorgadoEnderecoRua],
				Num:          extracted.Result[extractor.OutorgadoEnderecoN],
				CityUF:       extracted.Result[extractor.OutorgadoEnderecoCidadeUF],
				Neighborhood: extracted.Result[extractor.OutorgadoEnderecoBairro],
			},
		},
		InitialBookPages: extracted.Result[extractor.InitialBookPages],
		FinalBookPages:   extracted.Result[extractor.FinalBookPages],
	}

	return minuta.Minuta(params)
}

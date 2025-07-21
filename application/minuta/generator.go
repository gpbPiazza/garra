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

type GeneratorResponse struct {
	MinutaHTML     string   `json:"minuta_html"`
	TokensNotFound []string `json:"tokens_not_found"`
}

func (app *GeneratorApp) Generate(generateParams GenerateParams) (GeneratorResponse, error) {
	// ex := extractor.New()
	// ex.Extract(generateParams.DocStr)
	// extracted := ex.Result()

	ex2 := extractor.New2()
	ex2.Extract2(generateParams.DocStr)
	extracted := ex2.Result2()

	params := minuta.MinutaParams{
		TitleAto:            extracted.Scripture.TitleAto,
		TabelionatoName:     extracted.Scripture.Tablionato.Name,
		TabelionatoCityUF:   extracted.Scripture.Tablionato.CityUF,
		BookNum:             extracted.Scripture.BookNum,
		EscrituraMadeDate:   extracted.Scripture.EscrituraMadeDate,
		EscrituraValor:      extracted.Scripture.EscrituraValor,
		ItbiValor:           extracted.Scripture.ItbiValor,
		ItbiIncidenciaValor: extracted.Scripture.ItbiIncidenciaValor,
		Transmitente: minuta.PersonParams{
			IsOverqualified: generateParams.IsTransmitenteOverqualified,
			Name:            extracted.Scripture.Outorgantes[0].Name,
			Nationality:     extracted.Scripture.Outorgantes[0].Nationality,
			MaritalStatus:   extracted.Scripture.Outorgantes[0].MaritalStatus,
			DocNum_CPF_CNPJ: extracted.Scripture.Outorgantes[0].DocNum_CPF_CNPJ,
			DocType:         extracted.Scripture.Outorgantes[0].DocType,
			Sex:             extracted.Scripture.Outorgantes[0].Sex,
			Job:             extracted.Scripture.Outorgantes[0].Job,
			Address: minuta.AddressParams{
				Rua:          extracted.Scripture.Outorgantes[0].Address.Street,
				Num:          extracted.Scripture.Outorgantes[0].Address.Num,
				CityUF:       extracted.Scripture.Outorgantes[0].Address.CityUF,
				Neighborhood: extracted.Scripture.Outorgantes[0].Address.Neighborhood,
			},
		},
		Adquirente: minuta.PersonParams{
			IsOverqualified: generateParams.IsAdquirenteOverqualified,
			Name:            extracted.Scripture.Outorgados[0].Name,
			Nationality:     extracted.Scripture.Outorgados[0].Nationality,
			MaritalStatus:   extracted.Scripture.Outorgados[0].MaritalStatus,
			DocNum_CPF_CNPJ: extracted.Scripture.Outorgados[0].DocNum_CPF_CNPJ,
			DocType:         extracted.Scripture.Outorgados[0].DocType,
			Sex:             extracted.Scripture.Outorgados[0].Sex,
			Job:             extracted.Scripture.Outorgados[0].Job,
			Address: minuta.AddressParams{
				Rua:          extracted.Scripture.Outorgados[0].Address.Street,
				Num:          extracted.Scripture.Outorgados[0].Address.Num,
				CityUF:       extracted.Scripture.Outorgados[0].Address.CityUF,
				Neighborhood: extracted.Scripture.Outorgados[0].Address.Neighborhood,
			},
		},
		InitialBookPages: extracted.Scripture.TitleAto,
		FinalBookPages:   extracted.Scripture.TitleAto,
	}

	// params := minuta.MinutaParams{
	// 	TitleAto:            extracted.Result[extractor.TitleAto],
	// 	TabelionatoName:     extracted.Result[extractor.TabelionatoName],
	// 	TabelionatoCityUF:   extracted.Result[extractor.TabelionatoCityUF],
	// 	BookNum:             extracted.Result[extractor.BookNum],
	// 	EscrituraMadeDate:   extracted.Result[extractor.EscrituraMadeDate],
	// 	EscrituraValor:      extracted.Result[extractor.EscrituraValor],
	// 	ItbiValor:           extracted.Result[extractor.ItbiValor],
	// 	ItbiIncidenciaValor: extracted.Result[extractor.ItbiIncidenciaValor],
	// 	Transmitente: minuta.PersonParams{
	// 		IsOverqualified: generateParams.IsTransmitenteOverqualified,
	// 		Name:            extracted.Result[extractor.OutorganteName],
	// 		Nationality:     extracted.Result[extractor.OutorganteNationality],
	// 		MaritalStatus:   extracted.Result[extractor.OutorganteEstadoCivil],
	// 		DocNum_CPF_CNPJ: extracted.Result[extractor.OutorganteDocNumCPF_CNPJ],
	// 		DocType:         extracted.Result[extractor.OutorganteDocType],
	// 		Sex:             extracted.Result[extractor.OutorganteSex],
	// 		Job:             extracted.Result[extractor.OutorganteJob],
	// 		Address: minuta.AddressParams{
	// 			Rua:          extracted.Result[extractor.OutorganteEnderecoRua],
	// 			Num:          extracted.Result[extractor.OutorganteEnderecoN],
	// 			CityUF:       extracted.Result[extractor.OutorganteEnderecoCidadeUF],
	// 			Neighborhood: extracted.Result[extractor.OutorganteEnderecoBairro],
	// 		},
	// 	},
	// 	Adquirente: minuta.PersonParams{
	// 		IsOverqualified: generateParams.IsAdquirenteOverqualified,
	// 		Name:            extracted.Result[extractor.OutorgadoName],
	// 		Nationality:     extracted.Result[extractor.OutorgadoNationality],
	// 		MaritalStatus:   extracted.Result[extractor.OutorgadoEstadoCivil],
	// 		DocNum_CPF_CNPJ: extracted.Result[extractor.OutorgadoDocNumCPF_CNPJ],
	// 		DocType:         extracted.Result[extractor.OutorgadoDocType],
	// 		Sex:             extracted.Result[extractor.OutorgadoSex],
	// 		Job:             extracted.Result[extractor.OutorgadoJob],
	// 		Address: minuta.AddressParams{
	// 			Rua:          extracted.Result[extractor.OutorgadoEnderecoRua],
	// 			Num:          extracted.Result[extractor.OutorgadoEnderecoN],
	// 			CityUF:       extracted.Result[extractor.OutorgadoEnderecoCidadeUF],
	// 			Neighborhood: extracted.Result[extractor.OutorgadoEnderecoBairro],
	// 		},
	// 	},
	// 	InitialBookPages: extracted.Result[extractor.InitialBookPages],
	// 	FinalBookPages:   extracted.Result[extractor.FinalBookPages],
	// }

	minutaHTML, err := minuta.Minuta(params)
	if err != nil {
		return GeneratorResponse{}, err
	}

	return GeneratorResponse{
		MinutaHTML:     minutaHTML,
		TokensNotFound: app.mapTokensNotFound(extracted.TokensNotFound, params),
	}, nil
}

func (app *GeneratorApp) mapTokensNotFound(tokens []*extractor.Token, params minuta.MinutaParams) []string {
	var result []string

	for _, t := range tokens {
		if t.Identifier == extractor.OutorgadoJob && minuta.IsJuridicPerson(params.Adquirente.DocType) ||
			t.Identifier == extractor.OutorganteJob && minuta.IsJuridicPerson(params.Transmitente.DocType) {
			continue
		}

		result = append(result, extractor.IdentifiersNames[t.Identifier])
	}

	return result
}

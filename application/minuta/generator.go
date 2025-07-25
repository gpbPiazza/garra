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
			Job:             extracted.Result[extractor.OutorganteJob],
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
			Job:             extracted.Result[extractor.OutorgadoJob],
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

	minutaHTML, err := minuta.Minuta(params)
	if err != nil {
		return GeneratorResponse{}, err
	}

	var tokensNotFoundIdentifiers []string
	for _, t := range extracted.TokensNotFound {
		if t.Identifier == extractor.OutorgadoJob && minuta.IsJuridicPerson(params.Adquirente.DocType) ||
			t.Identifier == extractor.OutorganteJob && minuta.IsJuridicPerson(params.Transmitente.DocType) {
			continue
		}
		tokensNotFoundIdentifiers = append(tokensNotFoundIdentifiers, extractor.IdentifiersNames[t.Identifier])
	}

	return GeneratorResponse{
		MinutaHTML:     minutaHTML,
		TokensNotFound: tokensNotFoundIdentifiers,
	}, nil
}

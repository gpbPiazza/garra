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
	DocStr               string
	IsTransOverqualified bool
	IsAdquiOverqualified bool
}

type GeneratorResponse struct {
	MinutaHTML string `json:"minuta_html"`
}

func (app *GeneratorApp) Generate(params GenerateParams) (GeneratorResponse, error) {
	ex2 := extractor.New2()
	ex2.Extract2(params.DocStr)
	extracted := ex2.Result2()

	minutaHTML, err := minuta.Minuta(
		extracted.Scripture,
		params.IsTransOverqualified,
		params.IsAdquiOverqualified,
	)
	if err != nil {
		return GeneratorResponse{}, err
	}

	return GeneratorResponse{
		MinutaHTML: minutaHTML,
	}, nil
}

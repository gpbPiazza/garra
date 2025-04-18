package minuta

import (
	"errors"
	"strings"
)

// token is the relation between the data to find in a document and
// the Key to replace in a minuta template.
//
// Uses start and end keys to find Value
// Uses Replace to change Key into template by token.Value
type token struct {
	Start            string
	End              string
	Replace          ReplaceKey
	Value            string
	AlreadyExtracted bool
}

type Extractor struct {
	result map[ReplaceKey]string
	tokens []token
}

func NewExtractor() *Extractor {
	return &Extractor{
		result: make(map[ReplaceKey]string),
		tokens: []token{
			{
				Start:            "Registro Geral: MATRÍCULA Nº",
				End:              ",",
				Replace:          Matricula,
				AlreadyExtracted: false,
			},
			{
				Start:            "", // TODO: VER COM O ALEMÃO
				End:              "", // TODO: VER COM O ALEMÃO
				Replace:          TypeAto,
				AlreadyExtracted: false,
			},
			{
				Start:            "", // TODO: VER COM O ALEMÃO
				End:              "", // TODO: VER COM O ALEMÃO
				Replace:          NumAto,
				AlreadyExtracted: false,
			},
			{
				Start:            "Cláusula Geral: ESCRITURA PÚBLICA DE",
				End:              " que",
				Replace:          TitleAto,
				AlreadyExtracted: false,
			},
			{
				Start:            "",
				End:              "",
				Replace:          DataRegistro, // TODO: VER COM O ALEMÃO
				AlreadyExtracted: false,
			},
			{
				Start:            "",
				End:              "",
				Replace:          Protocolo, // TODO: VER COM O ALEMÃO
				AlreadyExtracted: false,
			},
			{
				Start:            "Data e hora do recebimento do ato pelo TJSC: ",
				End:              " -",
				Replace:          DataProtocolo,
				AlreadyExtracted: false,
			},
			{
				Start:            "Transmitente: ",
				End:              ".",
				Replace:          Transmitente, // igual a outorgante
				AlreadyExtracted: false,
			},
			{
				Start:            "Adquirente: ",
				End:              ".",
				Replace:          Adquirente, // igual a outorgado
				AlreadyExtracted: false,
			},
			{
				Start:            "Tabelionato Número: ",
				End:              ",",
				Replace:          TabelionatoNum,
				AlreadyExtracted: false,
			},
			{
				Start:            "Tabelionato Nome: ",
				End:              ",",
				Replace:          TabelionatoName,
				AlreadyExtracted: false,
			},
			{
				Start:            "Tabelionato Cidade/Estado: ",
				End:              ",",
				Replace:          TabelionatoCityState,
				AlreadyExtracted: false,
			},
			{
				Start:            "Livro Número: ",
				End:              ",",
				Replace:          BookNum,
				AlreadyExtracted: false,
			},
			{
				Start:            "Páginas do Livro: ",
				End:              ",",
				Replace:          BookPages,
				AlreadyExtracted: false,
			},
			{
				Start:            "Data da Escritura: ",
				End:              ",",
				Replace:          EscrituraMadeDate,
				AlreadyExtracted: false,
			},
			{
				Start:            "Valor da Escritura: R$ ",
				End:              ",",
				Replace:          EscrituraValor,
				AlreadyExtracted: false,
			},
			{
				Start:            "Valor Extenso da Escritura: ",
				End:              ".",
				Replace:          EscrituraValorExtenso,
				AlreadyExtracted: false,
			},
			{
				Start:            "Valor do ITBI: R$ ",
				End:              ",",
				Replace:          ItbiValor,
				AlreadyExtracted: false,
			},
			{
				Start:            "Valor de Incidência do ITBI: R$ ",
				End:              ",",
				Replace:          ItbiIncidenciaValor,
				AlreadyExtracted: false,
			},
		},
	}
}

func (e *Extractor) Extract(text string) {
	for _, token := range e.tokens {
		if token.AlreadyExtracted {
			continue
		}

		val, err := extractTokenValue(text, token)
		if err != nil {
			continue
		}
		token.Value = val
		token.AlreadyExtracted = true
		e.result[token.Replace] = token.Value
	}
}

func (e *Extractor) Result() map[ReplaceKey]string {
	return e.result
}

func extractTokenValue(text string, token token) (string, error) {
	startIndex := strings.Index(text, token.Start)
	if startIndex == -1 {
		return "", errors.New("token start index not found")
	}

	startIndex += len(token.Start)

	endIndex := strings.Index(text[startIndex:], token.End)
	if endIndex == -1 {
		return "", errors.New("token end index not found")
	}

	endIndex += startIndex

	return strings.TrimSpace(text[startIndex:endIndex]), nil
}

package extractor

import (
	"errors"
	"fmt"
	"log"
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
	Offset           string
	ResultKey        ResultKey
	Value            string
	AlreadyExtracted bool
}

type Extractor struct {
	result map[ResultKey]string
	tokens []*token
}

func New() *Extractor {
	return &Extractor{
		result: make(map[ResultKey]string),
		tokens: []*token{
			{
				Start:            "MATRÍCULA Nº",
				End:              ", CNM:",
				ResultKey:        Matricula,
				AlreadyExtracted: false,
			},
			// {
			// 	Start:            "", // TODO: VER COM O ALEMÃO
			// 	End:              "", // TODO: VER COM O ALEMÃO
			// 	ResultKey:        TypeAto,
			// 	AlreadyExtracted: false,
			// },
			// {
			// 	Start:            "", // TODO: VER COM O ALEMÃO
			// 	End:              "", // TODO: VER COM O ALEMÃO
			// 	ResultKey:        NumAto,
			// 	AlreadyExtracted: false,
			// },
			{
				Start:            "Cláusula Geral: ESCRITURA PÚBLICA DE",
				End:              " que",
				ResultKey:        TitleAto,
				AlreadyExtracted: false,
			},
			// {
			// 	Start:            "",
			// 	End:              "",
			// 	ResultKey:        DataRegistro, // TODO: VER COM O ALEMÃO
			// 	AlreadyExtracted: false,
			// },
			// {
			// 	Start:            "",
			// 	End:              "",
			// 	ResultKey:        Protocolo, // TODO: VER COM O ALEMÃO
			// 	AlreadyExtracted: false,
			// },
			{
				Start:            "Data e hora do recebimento do ato pelo TJSC: ",
				End:              " -",
				ResultKey:        DataProtocolo,
				AlreadyExtracted: false,
			},
			{
				Start:            "Parte :",
				End:              "Pessoa:",
				Offset:           "Outorgante",
				ResultKey:        OutorganteName,
				AlreadyExtracted: false,
			},
			{
				Start:            "Profissão: ",
				End:              " -",
				Offset:           "Outorgante",
				ResultKey:        OutorganteJob,
				AlreadyExtracted: false,
			},
			{
				Start:            "Nacionalidade: ",
				End:              " -",
				Offset:           "Outorgante",
				ResultKey:        OutorganteNationality,
				AlreadyExtracted: false,
			},
			{
				Start:            "Estado Civil: ",
				End:              "- Profissão:",
				Offset:           "Outorgante",
				ResultKey:        OutorganteEstadoCivil,
				AlreadyExtracted: false,
			},
			{
				Start:            "Doc. Nº:",
				End:              "/",
				Offset:           "Outorgante", // Esse cara ta errado ta pegando 2 caracteres não necssários
				ResultKey:        OutorganteDocNumCPF_CNPJ,
				AlreadyExtracted: false,
			},
			{
				Start:            "Doc. Tipo:",
				End:              "Doc.",
				Offset:           "Outorgante",
				ResultKey:        OutorganteDocType,
				AlreadyExtracted: false,
			},
			{
				Start:            "EndereçosLogradouro:",
				End:              "Número:",
				Offset:           "Outorgante",
				ResultKey:        OutorganteEnderecoRua,
				AlreadyExtracted: false,
			},
			{
				Start:            "Número: ",
				End:              "Bairro:",
				Offset:           "Outorgante",
				ResultKey:        OutorganteEnderecoN,
				AlreadyExtracted: false,
			},
			{
				Start:            "Bairro:",
				End:              "Complemento",
				Offset:           "Outorgante",
				ResultKey:        OutorganteEnderecoBairro,
				AlreadyExtracted: false,
			},
			{
				Start:            "Cidade/UF: ",
				End:              "CEP:",
				Offset:           "Outorgante",
				ResultKey:        OutorganteEnderecoCidadeUF,
				AlreadyExtracted: false,
			},
			{
				Start:            "Parte :",
				End:              "Data",
				Offset:           "Outorgado",
				ResultKey:        OutorgadoName,
				AlreadyExtracted: false,
			},
			{
				Start:            "Profissão:",
				End:              "- Nacionalidade:",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoJob,
				AlreadyExtracted: false,
			},
			{
				Start:            "Nacionalidade:",
				End:              "- Sexo:",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoNationality,
				AlreadyExtracted: false,
			},
			{
				Start:            "Estado Civil:",
				End:              "- Profissão:",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoEstadoCivil,
				AlreadyExtracted: false,
			},
			{
				Start:            "Doc. Nº:",
				End:              "Doc. Tipo:",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoDocNumCPF_CNPJ,
				AlreadyExtracted: false,
			},
			{
				Start:            "DocumentosDoc. Tipo: ",
				End:              "Doc.",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoDocType,
				AlreadyExtracted: false,
			},
			{
				Start:            "EndereçosLogradouro:",
				End:              "Número:",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoEnderecoRua,
				AlreadyExtracted: false,
			},
			{
				Start:            "Número:",
				End:              "Bairro:",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoEnderecoN,
				AlreadyExtracted: false,
			},
			{
				Start:            "Bairro: ",
				End:              "Cidade/UF:",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoEnderecoBairro,
				AlreadyExtracted: false,
			},
			{
				Start:            "Cidade/UF:",
				End:              "CEP:",
				Offset:           "OutorgadoParte",
				ResultKey:        OutorgadoEnderecoCidadeUF,
				AlreadyExtracted: false,
			},
			{
				Start:            " ",
				End:              "Endereço:",
				Offset:           "Serventia:",
				ResultKey:        TabelionatoName,
				AlreadyExtracted: false,
			},
			{
				Start:            "Município/UF:",
				End:              "Telefone(s):",
				Offset:           "Serventia:",
				ResultKey:        TabelionatoCityUF,
				AlreadyExtracted: false,
			},
			{
				Start:            "Livro: ",
				End:              "Nome do Livro:",
				Offset:           "RegistroCódigo do",
				ResultKey:        BookNum,
				AlreadyExtracted: false,
			},
			{
				Start:            "Página Inicial:",
				End:              "Página Final:",
				Offset:           "RegistroCódigo do",
				ResultKey:        InitialBookPages,
				AlreadyExtracted: false,
			},
			{
				Start:            "Página Final:",
				End:              "Data do Registro:",
				Offset:           "RegistroCódigo do",
				ResultKey:        FinalBookPages,
				AlreadyExtracted: false,
			},
			{
				Start:            "Data do Registro:",
				End:              "Nome do Imposto",
				Offset:           "RegistroCódigo do",
				ResultKey:        EscrituraMadeDate,
				AlreadyExtracted: false,
			},
			{
				Start:            "preço total, certo e ajustado de R$",
				End:              ", ",
				Offset:           "Cláusula Geral:",
				ResultKey:        EscrituraValor,
				AlreadyExtracted: false,
			},
			{
				Start:            "importância de R$",
				End:              " correspondente",
				Offset:           "Cláusula Geral:",
				ResultKey:        ItbiValor,
				AlreadyExtracted: false,
			},
			{
				Start:            "Valor do Negócio: R$",
				End:              "Cláusula Geral:",
				ResultKey:        ItbiIncidenciaValor,
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

		val, err := extractTokenValue(text, *token)
		if err != nil {
			log.Printf("extract token val err - token: '%s' - err: '%s'", resultKeyNames[token.ResultKey], err)
			continue
		}

		token.Value = val
		token.AlreadyExtracted = true
		e.result[token.ResultKey] = token.Value
	}
}

func (e *Extractor) Result() map[ResultKey]string {
	for _, t := range e.tokens {
		if !t.AlreadyExtracted {
			log.Printf("token not found - token: '%s'", resultKeyNames[t.ResultKey])
		}
	}

	return e.result
}

func extractTokenValue(text string, token token) (string, error) {
	if token.Start == "" {
		return "", errors.New("token start key is empty")
	}
	if token.End == "" {
		return "", errors.New("token end key is empty")
	}

	if token.Offset != "" {
		offSetIndex := strings.Index(text, token.Offset)
		if offSetIndex == -1 {
			return "", fmt.Errorf("offset '%s' not found in text", token.Offset)
		}
		text = text[offSetIndex+len(token.Offset):]
	}

	startIndex := strings.Index(text, token.Start)
	if startIndex == -1 {
		return "", fmt.Errorf("start key '%s' not found in text", token.Start)
	}

	startIndex += len(token.Start)

	endIndex := strings.Index(text[startIndex:], token.End)
	if endIndex == -1 {
		return "", fmt.Errorf("end key '%s' not found in text", token.End)
	}

	endIndex += startIndex

	return strings.TrimSpace(text[startIndex:endIndex]), nil
}

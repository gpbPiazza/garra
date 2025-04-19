package minuta

import (
	"errors"
	"fmt"
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
	NextEndOccurence bool
	Offset           string
	Replacer         ReplaceKey
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
				Replacer:         Matricula,
				AlreadyExtracted: false,
				Value:            "",
			},
			{
				Start:            "", // TODO: VER COM O ALEMÃO
				End:              "", // TODO: VER COM O ALEMÃO
				Replacer:         TypeAto,
				AlreadyExtracted: false,
			},
			{
				Start:            "", // TODO: VER COM O ALEMÃO
				End:              "", // TODO: VER COM O ALEMÃO
				Replacer:         NumAto,
				AlreadyExtracted: false,
			},
			{
				Start:            "Cláusula Geral: ESCRITURA PÚBLICA DE",
				End:              " que",
				Replacer:         TitleAto,
				AlreadyExtracted: false,
			},
			{
				Start:            "",
				End:              "",
				Replacer:         DataRegistro, // TODO: VER COM O ALEMÃO
				AlreadyExtracted: false,
			},
			{
				Start:            "",
				End:              "",
				Replacer:         Protocolo, // TODO: VER COM O ALEMÃO
				AlreadyExtracted: false,
			},
			{
				Start:            "Data e hora do recebimento do ato pelo TJSC: ",
				End:              " -",
				Replacer:         DataProtocolo,
				AlreadyExtracted: false,
			},
			{
				Start:            "Parte :",
				End:              "Pessoa:",
				Offset:           "Outorgante",
				Replacer:         TransmitenteNome,
				AlreadyExtracted: false,
			},
			{
				Start:            "Profissão: ",
				End:              " -",
				Offset:           "Outorgante",
				Replacer:         TransmitenteJob,
				AlreadyExtracted: false,
			},
			{
				Start:            "Nacionalidade: ",
				End:              " -",
				Offset:           "Outorgante",
				Replacer:         TransmitenteNacionalidade,
				AlreadyExtracted: false,
			},
			{
				Start:            "Estado Civil: ",
				End:              ",",
				Offset:           "Outorgante",
				Replacer:         TransmitenteEstadoCivil,
				AlreadyExtracted: false,
			},
			{
				Start:            "Doc. Nº:",
				End:              "/",
				Offset:           "Outorgante",
				Replacer:         TransmitenteCPF_CNPJ,
				AlreadyExtracted: false,
			},
			{
				Start:            "EndereçosLogradouro:",
				End:              "Número:",
				Offset:           "Outorgante",
				Replacer:         TransmitenteEnderecoRua,
				AlreadyExtracted: false,
			},
			{
				Start:            "Número: ",
				End:              "Bairro:",
				Offset:           "Outorgante",
				Replacer:         TransmitenteEnderecoN,
				AlreadyExtracted: false,
			},
			{
				Start:            "Bairro:",
				End:              "Complemento",
				Offset:           "Outorgante",
				Replacer:         TransmitenteEnderecoBairro,
				AlreadyExtracted: false,
			},
			{
				Start:            "Cidade/UF",
				End:              "CEP",
				Offset:           "Outorgante",
				Replacer:         TransmitenteEnderecoCidadeUF,
				AlreadyExtracted: false,
			},
			{
				Start:            "OutorgadoParte :",
				End:              "Data",
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteNome,
				AlreadyExtracted: false,
			},
			{
				Start:            "Profissão:",
				End:              "- Nacionalidade:",
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteJob,
				AlreadyExtracted: false,
			},
			{
				Start:            "Nacionalidade:",
				End:              "- Sexo:",
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteNacionalidade,
				AlreadyExtracted: false,
			},
			{
				Start:            "Estado Civil:",
				End:              "- Profissão:",
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteEstadoCivil,
				AlreadyExtracted: false,
			},
			{
				Start:            "Doc. Nº:",
				End:              "Doc. Tipo:",
				NextEndOccurence: true,
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteCPF_CNPJ,
				AlreadyExtracted: false,
			},
			{
				Start:            "EndereçosLogradouro:",
				End:              "Número:",
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteEnderecoRua,
				AlreadyExtracted: false,
			},
			{
				Start:            "Número:",
				End:              "Bairro:",
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteEnderecoN,
				AlreadyExtracted: false,
			},
			{
				Start:            "Bairro: ",
				End:              "Cidade/UF:",
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteEnderecoBairro,
				AlreadyExtracted: false,
			},
			{
				Start:            "Cidade/UF:",
				End:              "CEP:",
				Offset:           "OutorgadoParte",
				Replacer:         AdquirenteEnderecoCidadeUF,
				AlreadyExtracted: false,
			},
			{
				Start:            "Serventia: ",
				End:              "Endereço:",
				Offset:           "Serventia:",
				Replacer:         TabelionatoName,
				AlreadyExtracted: false,
			},
			{
				Start:            "Município/UF:",
				End:              "Telefone(s):",
				Offset:           "Serventia:",
				Replacer:         TabelionatoCityState,
				AlreadyExtracted: false,
			},
			{
				Start:            "Código do Livro: ",
				End:              "Nome do Livro:",
				Offset:           "RegistroCódigo do Livro:",
				Replacer:         BookNum,
				AlreadyExtracted: false,
			},
			{
				Start:            "Página Inicial:",
				End:              "Página Final:",
				Offset:           "RegistroCódigo do Livro:",
				Replacer:         InitialBookPage,
				AlreadyExtracted: false,
			},
			{
				Start:            "Página Final:",
				End:              "Data do Registro:",
				Offset:           "RegistroCódigo do Livro:",
				Replacer:         FinalBookPage,
				AlreadyExtracted: false,
			},
			{
				Start:            "Data do Registro:",
				End:              "Nome do Imposto",
				Offset:           "RegistroCódigo do Livro:",
				Replacer:         EscrituraMadeDate,
				AlreadyExtracted: false,
			},
			{
				Start:            "preço total, certo e ajustado de R$",
				End:              ", ",
				Offset:           "Cláusula Geral:",
				Replacer:         EscrituraValor,
				AlreadyExtracted: false,
			},
			{
				Start:            "importância de R$",
				End:              " correspondente",
				Offset:           "Cláusula Geral:",
				Replacer:         ItbiValor,
				AlreadyExtracted: false,
			},
			{
				Start:            "Valor do Negócio: R$",
				End:              "Cláusula Geral:",
				Replacer:         ItbiIncidenciaValor,
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
		e.result[token.Replacer] = token.Value
	}
}

func (e *Extractor) Result() map[ReplaceKey]string {
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

	// TODO: IMAPLEMENT SECOND OCCURENCE OF EndToken

	endIndex := strings.Index(text[startIndex:], token.End)
	if endIndex == -1 {
		return "", fmt.Errorf("end key '%s' not found in text", token.End)
	}

	endIndex += startIndex

	return strings.TrimSpace(text[startIndex:endIndex]), nil
}

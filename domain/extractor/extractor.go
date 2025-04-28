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
	Start       string
	End         string
	SecondEnd   string // TODO: End should be []string and not create a second key...
	Offset      string
	ResultKey   ResultKey
	Value       string
	IsExtracted bool
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
				Start:       "MATRÍCULA Nº",
				End:         ", CNM",
				ResultKey:   Matricula,
				IsExtracted: false,
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
				Start:       "Cláusula Geral: ESCRITURA PÚBLICA DE",
				End:         " que",
				ResultKey:   TitleAto,
				IsExtracted: false,
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
				Start:       "Data e hora do recebimento do ato pelo TJSC: ",
				End:         " -",
				ResultKey:   DataProtocolo,
				IsExtracted: false,
			},
			{
				Start:       "Parte :",
				End:         "Pessoa:",
				SecondEnd:   "Data de Nascimento:",
				Offset:      "Outorgante",
				ResultKey:   OutorganteName,
				IsExtracted: false,
			},
			// {
			// 	Start:       "Profissão: ",
			// 	End:         " -",
			// 	Offset:      "Outorgante",
			// 	ResultKey:   OutorganteJob,
			// 	IsExtracted: false,
			// },
			{
				Start:       "Nacionalidade: ",
				End:         " -",
				Offset:      "Outorgante",
				ResultKey:   OutorganteNationality,
				IsExtracted: false,
			},
			{
				Start:       "Estado Civil: ",
				End:         " -",
				Offset:      "Outorgante",
				ResultKey:   OutorganteEstadoCivil,
				IsExtracted: false,
			},
			{
				Start:       "Doc. Nº:",
				End:         "/",
				SecondEnd:   "Doc. Tipo:",
				Offset:      "Outorgante", // Esse cara ta errado ta pegando 2 caracteres não necssários
				ResultKey:   OutorganteDocNumCPF_CNPJ,
				IsExtracted: false,
			},
			{
				Start:       "Doc. Tipo:",
				End:         "Doc.",
				Offset:      "Outorgante",
				ResultKey:   OutorganteDocType,
				IsExtracted: false,
			},
			{
				Start:       "EndereçosLogradouro:",
				End:         "Número:",
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoRua,
				IsExtracted: false,
			},
			{
				Start:       "Número: ",
				End:         "Bairro:",
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoN,
				IsExtracted: false,
			},
			{
				Start:       "Bairro:",
				End:         "Complemento",
				SecondEnd:   "Cidade/UF:",
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoBairro,
				IsExtracted: false,
			},
			{
				Start:       "Cidade/UF: ",
				End:         "CEP:",
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoCidadeUF,
				IsExtracted: false,
			},
			{
				Start:       "Parte :",
				End:         "Data de Nascimento:",
				SecondEnd:   "Pessoa:",
				Offset:      "Outorgado",
				ResultKey:   OutorgadoName,
				IsExtracted: false,
			},
			// {
			// 	Start:       "Profissão:",
			// 	End:         "- Nacionalidade:",
			// 	Offset:      "OutorgadoParte",
			// 	ResultKey:   OutorgadoJob,
			// 	IsExtracted: false,
			// },
			{
				Start:       "Nacionalidade:",
				End:         "- Sexo:",
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoNationality,
				IsExtracted: false,
			},
			{
				Start:       "Estado Civil:",
				End:         "- Profissão:",
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEstadoCivil,
				IsExtracted: false,
			},
			{
				Start:       "Doc. Nº:",
				End:         "Doc. Tipo:",
				SecondEnd:   "EndereçosLogradouro:",
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoDocNumCPF_CNPJ,
				IsExtracted: false,
			},
			{
				Start:       "DocumentosDoc. Tipo: ",
				End:         "Doc.",
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoDocType,
				IsExtracted: false,
			},
			{
				Start:       "EndereçosLogradouro:",
				End:         "Número:",
				Offset:      "OutorgadoParte", // Para casos de pessoa jurídica é endereço da empresa ou do representante físico?
				ResultKey:   OutorgadoEnderecoRua,
				IsExtracted: false,
			},
			{
				Start:       "Número:",
				End:         "Bairro:",
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoN,
				IsExtracted: false,
			},
			{
				Start:       "Bairro: ",
				End:         "Cidade/UF:",
				SecondEnd:   "CEP:",
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoBairro,
				IsExtracted: false,
			},
			{
				Start:       "Cidade/UF:",
				End:         "CEP:",
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoCidadeUF,
				IsExtracted: false,
			},
			{
				Start:       " ",
				End:         "Endereço:",
				Offset:      "Serventia:",
				ResultKey:   TabelionatoName,
				IsExtracted: false,
			},
			{
				Start:       "Município/UF:",
				End:         "Telefone(s):",
				Offset:      "Serventia:",
				ResultKey:   TabelionatoCityUF,
				IsExtracted: false,
			},
			{
				Start:       "Livro: ",
				End:         "Nome do Livro:",
				Offset:      "RegistroCódigo do",
				ResultKey:   BookNum,
				IsExtracted: false,
			},
			{
				Start:       "Página Inicial:",
				End:         "Página Final:",
				Offset:      "RegistroCódigo do",
				ResultKey:   InitialBookPages,
				IsExtracted: false,
			},
			{
				Start:       "Página Final:",
				End:         "Data do Registro:",
				Offset:      "RegistroCódigo do",
				ResultKey:   FinalBookPages,
				IsExtracted: false,
			},
			{
				Start:       "Data do Registro:",
				End:         "Nome do Imposto",
				Offset:      "RegistroCódigo do",
				ResultKey:   EscrituraMadeDate,
				IsExtracted: false,
			},
			{
				Start:       "preço total, certo e ajustado de R$",
				End:         ", ",
				Offset:      "Cláusula Geral:",
				ResultKey:   EscrituraValor,
				IsExtracted: false,
			},
			{
				Start:       "importância de R$",
				End:         " correspondente",
				Offset:      "Cláusula Geral:",
				ResultKey:   ItbiValor,
				IsExtracted: false,
			},
			{
				Start:       "Valor do Negócio: R$",
				End:         "Cláusula Geral:",
				ResultKey:   ItbiIncidenciaValor,
				IsExtracted: false,
			},
		},
	}
}

func (e *Extractor) Extract(text string) {
	for _, token := range e.tokens {
		if token.IsExtracted {
			continue
		}

		val, err := extractTokenValue(text, *token)
		if err != nil {
			log.Printf("extract token val err - token: %s - err: %s", resultKeyNames[token.ResultKey], err)
			continue
		}

		token.Value = val
		token.IsExtracted = true
		e.result[token.ResultKey] = token.Value
	}
}

func (e *Extractor) Result() map[ResultKey]string {
	for _, t := range e.tokens {
		if !t.IsExtracted {
			log.Printf("token not found - token: '%s'", resultKeyNames[t.ResultKey])
		}

		if len(t.Value) >= 55 {
			log.Printf("maybe token value is incorrect - token: '%s'", resultKeyNames[t.ResultKey])
			log.Printf("token value: '%s'", t.Value)
			// log.Printf("token: '%+v'", t)
		}

		// if t.ResultKey == OutorgadoEnderecoBairro {
		// 	log.Printf("debug - token: '%s'", resultKeyNames[t.ResultKey])
		// 	log.Printf("token value: '%s'", t.Value)
		// }
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

	if token.SecondEnd != "" {
		secondEndIndex := strings.Index(text[startIndex:], token.SecondEnd)
		if secondEndIndex == -1 {
			return "", fmt.Errorf("second end key '%s' not found in text", token.SecondEnd)
		}

		if secondEndIndex < endIndex {
			endIndex = secondEndIndex
		}
	}

	endIndex += startIndex

	return strings.TrimSpace(text[startIndex:endIndex]), nil
}

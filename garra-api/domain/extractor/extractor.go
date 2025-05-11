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
	End         []string
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
				End:         []string{", CNM", ",CNM:"},
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
				End:         []string{" que"},
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
				End:         []string{" -"},
				ResultKey:   DataProtocolo,
				IsExtracted: false,
			},
			{
				Start:       "Parte :",
				End:         []string{"Pessoa:", "Data de Nascimento:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteName,
				IsExtracted: false,
			},
			// {
			// 	Start:       "Profissão: ",
			// 	End:         []string{" -"},
			// 	Offset:      "Outorgante",
			// 	ResultKey:   OutorganteJob,
			// 	IsExtracted: false,
			// },
			{
				Start:       "Nacionalidade: ",
				End:         []string{" -"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteNationality,
				IsExtracted: false,
			},
			{
				Start:       "Estado Civil: ",
				End:         []string{" -"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEstadoCivil,
				IsExtracted: false,
			},
			{
				Start:       "Sexo:",
				End:         []string{"DocumentosDoc."},
				Offset:      "Outorgante",
				ResultKey:   OutorganteSex,
				IsExtracted: false,
			},
			{
				Start:       "Doc. Nº:",
				End:         []string{"/", "Doc. Tipo:", "EndereçosLogradouro:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteDocNumCPF_CNPJ,
				IsExtracted: false,
			},
			{
				Start:       "Doc. Tipo:",
				End:         []string{"Doc."},
				Offset:      "Outorgante",
				ResultKey:   OutorganteDocType,
				IsExtracted: false,
			},
			{
				Start:       "EndereçosLogradouro:",
				End:         []string{"Número:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoRua,
				IsExtracted: false,
			},
			{
				Start:       "Número: ",
				End:         []string{"Bairro:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoN,
				IsExtracted: false,
			},
			{
				Start:       "Bairro:",
				End:         []string{"Complemento", "Cidade/UF:", ",", "Cidade:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoBairro,
				IsExtracted: false,
			},
			{
				Start:       "Cidade/UF: ",
				End:         []string{"CEP:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoCidadeUF,
				IsExtracted: false,
			},
			{
				Start:       "Parte :",
				End:         []string{"Data de Nascimento:", "Pessoa:"},
				Offset:      "Outorgado",
				ResultKey:   OutorgadoName,
				IsExtracted: false,
			},
			// {
			// 	Start:       "Profissão:",
			// 	End:         []string{"- Nacionalidade:"},
			// 	Offset:      "OutorgadoParte",
			// 	ResultKey:   OutorgadoJob,
			// 	IsExtracted: false,
			// },
			{
				Start:       "Nacionalidade:",
				End:         []string{"- Sexo:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoNationality,
				IsExtracted: false,
			},
			{
				Start:       "Sexo:",
				End:         []string{"DocumentosDoc."},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoSex,
				IsExtracted: false,
			},
			{
				Start:       "Estado Civil:",
				End:         []string{" -"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEstadoCivil,
				IsExtracted: false,
			},
			{
				Start:       "Doc. Nº:",
				End:         []string{"Doc. Tipo:", "EndereçosLogradouro:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoDocNumCPF_CNPJ,
				IsExtracted: false,
			},
			{
				Start:       "DocumentosDoc. Tipo: ",
				End:         []string{"Doc."},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoDocType,
				IsExtracted: false,
			},
			{
				Start:       "EndereçosLogradouro:",
				End:         []string{"Número:"},
				Offset:      "OutorgadoParte", // Para casos de pessoa jurídica é endereço da empresa ou do representante físico?
				ResultKey:   OutorgadoEnderecoRua,
				IsExtracted: false,
			},
			{
				Start:       "Número:",
				End:         []string{"Bairro:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoN,
				IsExtracted: false,
			},
			{
				Start:       "Bairro: ",
				End:         []string{"Cidade/UF:", "CEP:", "Complemento"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoBairro,
				IsExtracted: false,
			},
			{
				Start:       "Cidade/UF:",
				End:         []string{"CEP:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoCidadeUF,
				IsExtracted: false,
			},
			{
				Start:       " ",
				End:         []string{"Endereço:"},
				Offset:      "Serventia:",
				ResultKey:   TabelionatoName,
				IsExtracted: false,
			},
			{
				Start:       "Município/UF:",
				End:         []string{"Telefone(s):"},
				Offset:      "Serventia:",
				ResultKey:   TabelionatoCityUF,
				IsExtracted: false,
			},
			{
				Start:       "Livro: ",
				End:         []string{"Nome do Livro:"},
				Offset:      "RegistroCódigo do",
				ResultKey:   BookNum,
				IsExtracted: false,
			},
			{
				Start:       "Página Inicial:",
				End:         []string{"Página Final:"},
				Offset:      "RegistroCódigo do",
				ResultKey:   InitialBookPages,
				IsExtracted: false,
			},
			{
				Start:       "Página Final:",
				End:         []string{"Data do Registro:"},
				Offset:      "RegistroCódigo do",
				ResultKey:   FinalBookPages,
				IsExtracted: false,
			},
			{
				Start:       "Data do Registro: ",
				End:         []string{"Nome do Imposto", ","},
				Offset:      "RegistroCódigo do Livro:",
				ResultKey:   EscrituraMadeDate,
				IsExtracted: false,
			},
			{
				Start:       "preço total, certo e ajustado de R$",
				End:         []string{", ", "/"},
				Offset:      "Cláusula Geral:",
				ResultKey:   EscrituraValor,
				IsExtracted: false,
			},
			{
				Start:       "importância de R$",
				End:         []string{" correspondente"},
				Offset:      "Cláusula Geral:",
				ResultKey:   ItbiValor,
				IsExtracted: false,
			},
			{
				Start:       "Valor do Negócio: R$",
				End:         []string{"Cláusula Geral:"},
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
	if len(token.End) == 0 {
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

	// if token.SecondEnd == "" {
	// 	return "", fmt.Errorf("end key '%s' not found in text", token.End)
	// }

	endIndex := -1
	for _, end := range token.End {
		currentEndIndex := strings.Index(text[startIndex:], end)
		if currentEndIndex == -1 {
			continue
		}

		isFirstIteration := endIndex == -1
		if isFirstIteration || currentEndIndex < endIndex {
			endIndex = currentEndIndex
		}
	}

	if endIndex == -1 {
		return "", fmt.Errorf("end key '%s' not found in text", token.End)
	}

	endIndex += startIndex

	return strings.TrimSpace(text[startIndex:endIndex]), nil
}

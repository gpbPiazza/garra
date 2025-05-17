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
	StartKeys   []string
	EndKeys     []string
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
				StartKeys:   []string{"MATRÍCULA Nº"},
				EndKeys:     []string{", CNM", ",CNM:"},
				ResultKey:   Matricula,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Cláusula Geral: ESCRITURA PÚBLICA DE"},
				EndKeys:     []string{" que"},
				ResultKey:   TitleAto,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Data e hora do recebimento do ato pelo TJSC: "},
				EndKeys:     []string{" -"},
				ResultKey:   DataProtocolo,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Parte :"},
				EndKeys:     []string{"Pessoa:", "Data de Nascimento:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteName,
				IsExtracted: false,
			},
			// {
			// 	Start:       []string{"Profissão: ",},
			// 	End:         []string{" -"},
			// 	Offset:      "Outorgante",
			// 	ResultKey:   OutorganteJob,
			// 	IsExtracted: false,
			// },
			{
				StartKeys:   []string{"Nacionalidade: "},
				EndKeys:     []string{" -"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteNationality,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Estado Civil: "},
				EndKeys:     []string{" -"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEstadoCivil,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Sexo:"},
				EndKeys:     []string{"DocumentosDoc."},
				Offset:      "Outorgante",
				ResultKey:   OutorganteSex,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Doc. Nº:"},
				EndKeys:     []string{"/", "Doc. Tipo:", "EndereçosLogradouro:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteDocNumCPF_CNPJ,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Doc. Tipo:"},
				EndKeys:     []string{"Doc."},
				Offset:      "Outorgante",
				ResultKey:   OutorganteDocType,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"EndereçosLogradouro:"},
				EndKeys:     []string{"Número:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoRua,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Número: "},
				EndKeys:     []string{"Bairro:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoN,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Bairro:"},
				EndKeys:     []string{"Complemento", "Cidade/UF:", ",", "Cidade:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoBairro,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Cidade/UF: "},
				EndKeys:     []string{"CEP:"},
				Offset:      "Outorgante",
				ResultKey:   OutorganteEnderecoCidadeUF,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Parte :"},
				EndKeys:     []string{"Data de Nascimento:", "Pessoa:"},
				Offset:      "Outorgado",
				ResultKey:   OutorgadoName,
				IsExtracted: false,
			},
			// {
			// 	Start:       []string{"Profissão:",},
			// 	End:         []string{"- Nacionalidade:"},
			// 	Offset:      "OutorgadoParte",
			// 	ResultKey:   OutorgadoJob,
			// 	IsExtracted: false,
			// },
			{
				StartKeys:   []string{"Nacionalidade:"},
				EndKeys:     []string{"- Sexo:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoNationality,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Sexo:"},
				EndKeys:     []string{"DocumentosDoc."},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoSex,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Estado Civil:"},
				EndKeys:     []string{" -"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEstadoCivil,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Doc. Nº:"},
				EndKeys:     []string{"Doc. Tipo:", "EndereçosLogradouro:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoDocNumCPF_CNPJ,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"DocumentosDoc. Tipo: "},
				EndKeys:     []string{"Doc."},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoDocType,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"EndereçosLogradouro:"},
				EndKeys:     []string{"Número:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoRua,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Número:"},
				EndKeys:     []string{"Bairro:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoN,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Bairro: "},
				EndKeys:     []string{"Cidade/UF:", "CEP:", "Complemento"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoBairro,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Cidade/UF:"},
				EndKeys:     []string{"CEP:"},
				Offset:      "OutorgadoParte",
				ResultKey:   OutorgadoEnderecoCidadeUF,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{" "},
				EndKeys:     []string{"Endereço:"},
				Offset:      "Serventia:",
				ResultKey:   TabelionatoName,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Município/UF:"},
				EndKeys:     []string{"Telefone(s):"},
				Offset:      "Serventia:",
				ResultKey:   TabelionatoCityUF,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{":"},
				EndKeys:     []string{"Nome do Livro:"},
				Offset:      "Código do Livro",
				ResultKey:   BookNum,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Página Inicial:"},
				EndKeys:     []string{"Página Final:"},
				Offset:      "Código do Livro",
				ResultKey:   InitialBookPages,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Página Final:"},
				EndKeys:     []string{"Data do Registro:"},
				Offset:      "Código do Livro",
				ResultKey:   FinalBookPages,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Data do Registro: "},
				EndKeys:     []string{"Nome do Imposto", ","},
				Offset:      "Código do Livro",
				ResultKey:   EscrituraMadeDate,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"preço total, certo e ajustado de R$", "certo e ajustadode R$"},
				EndKeys:     []string{", ", "/"},
				Offset:      "Cláusula Geral:",
				ResultKey:   EscrituraValor,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"importância de R$"},
				EndKeys:     []string{" correspondente"},
				Offset:      "Cláusula Geral:",
				ResultKey:   ItbiValor,
				IsExtracted: false,
			},
			{
				StartKeys:   []string{"Valor do Negócio: R$"},
				EndKeys:     []string{"Cláusula Geral:"},
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
	if len(token.StartKeys) == 0 {
		return "", errors.New("token start key is empty")
	}
	if len(token.EndKeys) == 0 {
		return "", errors.New("token end key is empty")
	}

	if token.Offset != "" {
		offSetIndex := strings.Index(text, token.Offset)
		if offSetIndex == -1 {
			return "", fmt.Errorf("offset '%s' not found in text", token.Offset)
		}
		text = text[offSetIndex+len(token.Offset):]
	}

	startIndex := -1
	for _, start := range token.StartKeys {
		currentStartIndex := strings.Index(text, start)
		if currentStartIndex == -1 {
			continue
		}

		if currentStartIndex > startIndex {
			startIndex = currentStartIndex
			startIndex += len(start)
		}
	}

	if startIndex == -1 {
		return "", errors.New("any of start keys found in text")
	}

	endIndex := -1
	for _, end := range token.EndKeys {
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
		return "", errors.New("any of end keys found in text")
	}

	endIndex += startIndex

	return strings.TrimSpace(text[startIndex:endIndex]), nil
}

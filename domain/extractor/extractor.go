package extractor

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Token is the relation between the data to find in a document and
// the Key to replace in a minuta template.
//
// Uses start and end keys to find Value
// Uses Replace to change Key into template by Token.Value
type Token struct {
	StartKeys, EndKeys     []string
	StartOffset, EndOffset string
	Identifier             Identifier
	Value                  string
	IsExtracted            bool
}

type Extractor struct {
	result map[Identifier]string
	tokens []*Token

	outorgadoPersons  []Person
	outorgantePersons []Person

	TokensNotFound []*Token

	result2 map[Identifier]*Token
}

func NewScriptureTokens() []*Token {
	return []*Token{
		{
			StartKeys:   []string{"MATRÍCULA Nº"},
			EndKeys:     []string{", CNM", ",CNM:"},
			Identifier:  Matricula,
			IsExtracted: false,
			Value:       defaultValue(Matricula),
		},
		{
			StartKeys:   []string{"Cláusula Geral: ESCRITURA PÚBLICA DE"},
			EndKeys:     []string{" que"},
			Identifier:  TitleAto,
			IsExtracted: false,
			Value:       defaultValue(TitleAto),
		},
		{
			StartKeys:   []string{"Data e hora do recebimento do ato pelo TJSC: "},
			EndKeys:     []string{" -"},
			Identifier:  DataProtocolo,
			IsExtracted: false,
			Value:       defaultValue(DataProtocolo),
		},
		{
			StartKeys:   []string{" "},
			EndKeys:     []string{"Endereço:"},
			StartOffset: "Serventia:",
			Identifier:  TabelionatoName,
			IsExtracted: false,
			Value:       defaultValue(TabelionatoName),
		},
		{
			StartKeys:   []string{"Município/UF:"},
			EndKeys:     []string{"Telefone(s):"},
			StartOffset: "Serventia:",
			Identifier:  TabelionatoCityUF,
			IsExtracted: false,
			Value:       defaultValue(TabelionatoCityUF),
		},
		{
			StartKeys:   []string{":"},
			EndKeys:     []string{"Nome do Livro:"},
			StartOffset: "Código do Livro",
			Identifier:  BookNum,
			IsExtracted: false,
			Value:       defaultValue(BookNum),
		},
		{
			StartKeys:   []string{"Página Inicial:"},
			EndKeys:     []string{"Página Final:"},
			StartOffset: "Código do Livro",
			Identifier:  InitialBookPages,
			IsExtracted: false,
			Value:       defaultValue(InitialBookPages),
		},
		{
			StartKeys:   []string{"Página Final:"},
			EndKeys:     []string{"Data do Registro:"},
			StartOffset: "Código do Livro",
			Identifier:  FinalBookPages,
			IsExtracted: false,
			Value:       defaultValue(FinalBookPages),
		},
		{
			StartKeys:   []string{"Data do Registro: "},
			EndKeys:     []string{"Nome do Imposto", ","},
			StartOffset: "Código do Livro",
			Identifier:  EscrituraMadeDate,
			IsExtracted: false,
			Value:       defaultValue(EscrituraMadeDate),
		},
		{
			StartKeys:   []string{"preço total, certo e ajustado de R$", "certo e ajustadode R$"},
			EndKeys:     []string{", ", "/"},
			StartOffset: "Cláusula Geral:",
			Identifier:  EscrituraValor,
			IsExtracted: false,
			Value:       defaultValue(EscrituraValor),
		},
		{
			StartKeys:   []string{"importância de R$"},
			EndKeys:     []string{" correspondente"},
			StartOffset: "Cláusula Geral:",
			Identifier:  ItbiValor,
			IsExtracted: false,
			Value:       defaultValue(ItbiValor),
		},
		{
			StartKeys:   []string{"Valor do Negócio: R$"},
			EndKeys:     []string{"Cláusula Geral:"},
			Identifier:  ItbiIncidenciaValor,
			IsExtracted: false,
			Value:       defaultValue(ItbiIncidenciaValor),
		},
	}
}

func NewOutorganteTokens() []*Token {
	return []*Token{
		{
			StartKeys:   []string{"Parte :"},
			EndKeys:     []string{"Pessoa:", "Data de Nascimento:"},
			StartOffset: "Outorgante",
			Identifier:  OutorganteName,
			IsExtracted: false,
			Value:       defaultValue(OutorganteName),
		},
		{
			StartKeys:   []string{"Profissão:"},
			EndKeys:     []string{" - Nacionalidade:"},
			StartOffset: "Outorgante",
			EndOffset:   "OutorgadoParte",
			Identifier:  OutorganteJob,
			IsExtracted: false,
			Value:       defaultValue(OutorganteJob),
		},
		{
			StartKeys:   []string{"Nacionalidade: "},
			EndKeys:     []string{" -"},
			StartOffset: "Outorgante",
			Identifier:  OutorganteNationality,
			IsExtracted: false,
			Value:       defaultValue(OutorganteNationality),
		},
		{
			StartKeys:   []string{"Estado Civil: "},
			EndKeys:     []string{" -"},
			StartOffset: "Outorgante",
			Identifier:  OutorganteEstadoCivil,
			IsExtracted: false,
			Value:       defaultValue(OutorganteEstadoCivil),
		},
		{
			StartKeys:   []string{"Sexo:"},
			EndKeys:     []string{"DocumentosDoc."},
			StartOffset: "Outorgante",
			Identifier:  OutorganteSex,
			IsExtracted: false,
			Value:       defaultValue(OutorganteSex),
		},
		{
			StartKeys:   []string{"Doc. Nº:"},
			EndKeys:     []string{"/", "Doc. Tipo:", "EndereçosLogradouro:"},
			StartOffset: "Outorgante",
			Identifier:  OutorganteDocNumCPF_CNPJ,
			IsExtracted: false,
			Value:       defaultValue(OutorganteDocNumCPF_CNPJ),
		},
		{
			StartKeys:   []string{"Doc. Tipo:"},
			EndKeys:     []string{"Doc."},
			StartOffset: "Outorgante",
			Identifier:  OutorganteDocType,
			IsExtracted: false,
			Value:       defaultValue(OutorganteDocType),
		},
		{
			StartKeys:   []string{"EndereçosLogradouro:"},
			EndKeys:     []string{"Número:"},
			StartOffset: "Outorgante",
			Identifier:  OutorganteEnderecoRua,
			IsExtracted: false,
			Value:       defaultValue(OutorganteEnderecoRua),
		},
		{
			StartKeys:   []string{"Número: "},
			EndKeys:     []string{"Bairro:"},
			StartOffset: "Outorgante",
			Identifier:  OutorganteEnderecoN,
			IsExtracted: false,
			Value:       defaultValue(OutorganteEnderecoN),
		},
		{
			StartKeys:   []string{"Bairro:"},
			EndKeys:     []string{"Complemento", "Cidade/UF:", ",", "Cidade:"},
			StartOffset: "Outorgante",
			Identifier:  OutorganteEnderecoBairro,
			IsExtracted: false,
			Value:       defaultValue(OutorganteEnderecoBairro),
		},
		{
			StartKeys:   []string{"Cidade/UF: "},
			EndKeys:     []string{"CEP:"},
			StartOffset: "Outorgante",
			Identifier:  OutorganteEnderecoCidadeUF,
			IsExtracted: false,
			Value:       defaultValue(OutorganteEnderecoCidadeUF),
		},
	}
}

func NewOutorgadoTokens() []*Token {
	return []*Token{
		{
			StartKeys:   []string{"Parte :"},
			EndKeys:     []string{"Data de Nascimento:", "Pessoa:"},
			StartOffset: "Outorgado",
			Identifier:  OutorgadoName,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoName),
		},
		{
			StartKeys:   []string{"Profissão:"},
			EndKeys:     []string{"- Nacionalidade:"},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoJob,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoJob),
		},
		{
			StartKeys:   []string{"Nacionalidade:"},
			EndKeys:     []string{"- Sexo:"},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoNationality,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoNationality),
		},
		{
			StartKeys:   []string{"Sexo:"},
			EndKeys:     []string{"DocumentosDoc."},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoSex,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoSex),
		},
		{
			StartKeys:   []string{"Estado Civil:"},
			EndKeys:     []string{" -"},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoEstadoCivil,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoEstadoCivil),
		},
		{
			StartKeys:   []string{"Doc. Nº:"},
			EndKeys:     []string{"Doc. Tipo:", "EndereçosLogradouro:"},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoDocNumCPF_CNPJ,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoDocNumCPF_CNPJ),
		},
		{
			StartKeys:   []string{"DocumentosDoc. Tipo: "},
			EndKeys:     []string{"Doc."},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoDocType,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoDocType),
		},
		{
			StartKeys:   []string{"EndereçosLogradouro:"},
			EndKeys:     []string{"Número:"},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoEnderecoRua,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoEnderecoRua),
		},
		{
			StartKeys:   []string{"Número:"},
			EndKeys:     []string{"Bairro:"},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoEnderecoN,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoEnderecoN),
		},
		{
			StartKeys:   []string{"Bairro: "},
			EndKeys:     []string{"Cidade/UF:", "CEP:", "Complemento"},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoEnderecoBairro,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoEnderecoBairro),
		},
		{
			StartKeys:   []string{"Cidade/UF:"},
			EndKeys:     []string{"CEP:"},
			StartOffset: "OutorgadoParte",
			Identifier:  OutorgadoEnderecoCidadeUF,
			IsExtracted: false,
			Value:       defaultValue(OutorgadoEnderecoCidadeUF),
		},
	}
}

func New() *Extractor {
	var tokens []*Token
	tokens = append(tokens, NewScriptureTokens()...)
	tokens = append(tokens, NewOutorganteTokens()...)
	tokens = append(tokens, NewOutorgadoTokens()...)

	return &Extractor{
		result: make(map[Identifier]string),
		tokens: tokens,
	}
}

func (e *Extractor) Extract(text string) {
	for _, token := range e.tokens {
		if token.IsExtracted {
			continue
		}

		e.result[token.Identifier] = token.Value

		val, err := extractTokenValue(text, *token)
		if err != nil {
			log.Printf("extract token val err - token: %s - err: %s", IdentifiersNames[token.Identifier], err)
			continue
		}

		token.Value = val
		token.IsExtracted = true
		e.result[token.Identifier] = val
	}
}

func extractTokenValue(text string, token Token) (string, error) {
	if len(token.StartKeys) == 0 {
		return "", errors.New("token start key is empty")
	}
	if len(token.EndKeys) == 0 {
		return "", errors.New("token end key is empty")
	}

	text, err := cutOffSet(text, token.StartOffset, token.EndOffset)
	if err != nil {
		return "", err
	}

	return extractValue(text, token)
}

var NotFoundDefaultSuffix = "NÃO ENCONTRADO"

// defaultValue is just identifier name + not found string
func defaultValue(i Identifier) string {
	return fmt.Sprintf("[[%s %s]]", IdentifiersNames[i], NotFoundDefaultSuffix)
}

func cutOffSet(text, start, end string) (string, error) {
	if start != "" {
		offSetIndex := strings.Index(text, start)
		if offSetIndex == -1 {
			return "", fmt.Errorf("offset '%s' not found in text", start)
		}
		text = text[offSetIndex+len(start):]
	}

	if end != "" {
		endOffSetIndex := strings.Index(text, end)
		if endOffSetIndex == -1 {
			return "", fmt.Errorf("offset '%s' not found in text", end)
		}
		text = text[:endOffSetIndex+len(end)]
	}

	return text, nil
}

func extractValue(text string, token Token) (string, error) {
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

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
	StartKeys, EndKeys []string
	Offset, EndOffset  string
	Identifier         Identifier
	Value              string
	IsExtracted        bool
}

type Extractor struct {
	result map[Identifier]string
	tokens []*Token
}

func New() *Extractor {
	return &Extractor{
		result: make(map[Identifier]string),
		tokens: []*Token{
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
				StartKeys:   []string{"Parte :"},
				EndKeys:     []string{"Pessoa:", "Data de Nascimento:"},
				Offset:      "Outorgante",
				Identifier:  OutorganteName,
				IsExtracted: false,
				Value:       defaultValue(OutorganteName),
			},
			{
				StartKeys:   []string{"Profissão:"},
				EndKeys:     []string{" - Nacionalidade:"},
				Offset:      "Outorgante",
				EndOffset:   "OutorgadoParte",
				Identifier:  OutorganteJob,
				IsExtracted: false,
				Value:       defaultValue(OutorganteJob),
			},
			{
				StartKeys:   []string{"Nacionalidade: "},
				EndKeys:     []string{" -"},
				Offset:      "Outorgante",
				Identifier:  OutorganteNationality,
				IsExtracted: false,
				Value:       defaultValue(OutorganteNationality),
			},
			{
				StartKeys:   []string{"Estado Civil: "},
				EndKeys:     []string{" -"},
				Offset:      "Outorgante",
				Identifier:  OutorganteEstadoCivil,
				IsExtracted: false,
				Value:       defaultValue(OutorganteEstadoCivil),
			},
			{
				StartKeys:   []string{"Sexo:"},
				EndKeys:     []string{"DocumentosDoc."},
				Offset:      "Outorgante",
				Identifier:  OutorganteSex,
				IsExtracted: false,
				Value:       defaultValue(OutorganteSex),
			},
			{
				StartKeys:   []string{"Doc. Nº:"},
				EndKeys:     []string{"/", "Doc. Tipo:", "EndereçosLogradouro:"},
				Offset:      "Outorgante",
				Identifier:  OutorganteDocNumCPF_CNPJ,
				IsExtracted: false,
				Value:       defaultValue(OutorganteDocNumCPF_CNPJ),
			},
			{
				StartKeys:   []string{"Doc. Tipo:"},
				EndKeys:     []string{"Doc."},
				Offset:      "Outorgante",
				Identifier:  OutorganteDocType,
				IsExtracted: false,
				Value:       defaultValue(OutorganteDocType),
			},
			{
				StartKeys:   []string{"EndereçosLogradouro:"},
				EndKeys:     []string{"Número:"},
				Offset:      "Outorgante",
				Identifier:  OutorganteEnderecoRua,
				IsExtracted: false,
				Value:       defaultValue(OutorganteEnderecoRua),
			},
			{
				StartKeys:   []string{"Número: "},
				EndKeys:     []string{"Bairro:"},
				Offset:      "Outorgante",
				Identifier:  OutorganteEnderecoN,
				IsExtracted: false,
				Value:       defaultValue(OutorganteEnderecoN),
			},
			{
				StartKeys:   []string{"Bairro:"},
				EndKeys:     []string{"Complemento", "Cidade/UF:", ",", "Cidade:"},
				Offset:      "Outorgante",
				Identifier:  OutorganteEnderecoBairro,
				IsExtracted: false,
				Value:       defaultValue(OutorganteEnderecoBairro),
			},
			{
				StartKeys:   []string{"Cidade/UF: "},
				EndKeys:     []string{"CEP:"},
				Offset:      "Outorgante",
				Identifier:  OutorganteEnderecoCidadeUF,
				IsExtracted: false,
				Value:       defaultValue(OutorganteEnderecoCidadeUF),
			},
			{
				StartKeys:   []string{"Parte :"},
				EndKeys:     []string{"Data de Nascimento:", "Pessoa:"},
				Offset:      "Outorgado",
				Identifier:  OutorgadoName,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoName),
			},
			{
				StartKeys:   []string{"Profissão:"},
				EndKeys:     []string{"- Nacionalidade:"},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoJob,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoJob),
			},
			{
				StartKeys:   []string{"Nacionalidade:"},
				EndKeys:     []string{"- Sexo:"},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoNationality,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoNationality),
			},
			{
				StartKeys:   []string{"Sexo:"},
				EndKeys:     []string{"DocumentosDoc."},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoSex,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoSex),
			},
			{
				StartKeys:   []string{"Estado Civil:"},
				EndKeys:     []string{" -"},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoEstadoCivil,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoEstadoCivil),
			},
			{
				StartKeys:   []string{"Doc. Nº:"},
				EndKeys:     []string{"Doc. Tipo:", "EndereçosLogradouro:"},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoDocNumCPF_CNPJ,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoDocNumCPF_CNPJ),
			},
			{
				StartKeys:   []string{"DocumentosDoc. Tipo: "},
				EndKeys:     []string{"Doc."},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoDocType,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoDocType),
			},
			{
				StartKeys:   []string{"EndereçosLogradouro:"},
				EndKeys:     []string{"Número:"},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoEnderecoRua,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoEnderecoRua),
			},
			{
				StartKeys:   []string{"Número:"},
				EndKeys:     []string{"Bairro:"},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoEnderecoN,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoEnderecoN),
			},
			{
				StartKeys:   []string{"Bairro: "},
				EndKeys:     []string{"Cidade/UF:", "CEP:", "Complemento"},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoEnderecoBairro,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoEnderecoBairro),
			},
			{
				StartKeys:   []string{"Cidade/UF:"},
				EndKeys:     []string{"CEP:"},
				Offset:      "OutorgadoParte",
				Identifier:  OutorgadoEnderecoCidadeUF,
				IsExtracted: false,
				Value:       defaultValue(OutorgadoEnderecoCidadeUF),
			},
			{
				StartKeys:   []string{" "},
				EndKeys:     []string{"Endereço:"},
				Offset:      "Serventia:",
				Identifier:  TabelionatoName,
				IsExtracted: false,
				Value:       defaultValue(TabelionatoName),
			},
			{
				StartKeys:   []string{"Município/UF:"},
				EndKeys:     []string{"Telefone(s):"},
				Offset:      "Serventia:",
				Identifier:  TabelionatoCityUF,
				IsExtracted: false,
				Value:       defaultValue(TabelionatoCityUF),
			},
			{
				StartKeys:   []string{":"},
				EndKeys:     []string{"Nome do Livro:"},
				Offset:      "Código do Livro",
				Identifier:  BookNum,
				IsExtracted: false,
				Value:       defaultValue(BookNum),
			},
			{
				StartKeys:   []string{"Página Inicial:"},
				EndKeys:     []string{"Página Final:"},
				Offset:      "Código do Livro",
				Identifier:  InitialBookPages,
				IsExtracted: false,
				Value:       defaultValue(InitialBookPages),
			},
			{
				StartKeys:   []string{"Página Final:"},
				EndKeys:     []string{"Data do Registro:"},
				Offset:      "Código do Livro",
				Identifier:  FinalBookPages,
				IsExtracted: false,
				Value:       defaultValue(FinalBookPages),
			},
			{
				StartKeys:   []string{"Data do Registro: "},
				EndKeys:     []string{"Nome do Imposto", ","},
				Offset:      "Código do Livro",
				Identifier:  EscrituraMadeDate,
				IsExtracted: false,
				Value:       defaultValue(EscrituraMadeDate),
			},
			{
				StartKeys:   []string{"preço total, certo e ajustado de R$", "certo e ajustadode R$"},
				EndKeys:     []string{", ", "/"},
				Offset:      "Cláusula Geral:",
				Identifier:  EscrituraValor,
				IsExtracted: false,
				Value:       defaultValue(EscrituraValor),
			},
			{
				StartKeys:   []string{"importância de R$"},
				EndKeys:     []string{" correspondente"},
				Offset:      "Cláusula Geral:",
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
		},
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

	if token.Offset != "" {
		offSetIndex := strings.Index(text, token.Offset)
		if offSetIndex == -1 {
			return "", fmt.Errorf("offset '%s' not found in text", token.Offset)
		}
		text = text[offSetIndex+len(token.Offset):]
	}

	if token.EndOffset != "" {
		endOffSetIndex := strings.Index(text, token.EndOffset)
		if endOffSetIndex == -1 {
			return "", fmt.Errorf("offset '%s' not found in text", token.EndOffset)
		}
		text = text[:endOffSetIndex+len(token.EndOffset)]
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

var NotFoundDefaultSuffix = "NÃO ENCONTRADO"

// defaultValue is just identifier name + not found string
func defaultValue(i Identifier) string {
	return fmt.Sprintf("[[%s %s]]", IdentifiersNames[i], NotFoundDefaultSuffix)
}

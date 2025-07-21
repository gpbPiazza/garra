package extractor

import (
	"log"
	"strings"
)

type Address struct {
	Street       string
	Num          string
	CityUF       string
	Neighborhood string
}

type Person struct {
	Name            string
	Nationality     string
	Job             string
	MaritalStatus   string
	DocNum_CPF_CNPJ string
	DocType         string
	Address         Address
	Sex             string
}

type Tablionato struct {
	Name   string
	CityUF string
}

type Scripture struct {
	Outorgantes []Person
	Outorgados  []Person
	Tablionato  Tablionato

	TitleAto            string
	BookNum             string
	InitialBookPages    string
	FinalBookPages      string
	EscrituraMadeDate   string
	EscrituraValor      string
	ItbiValor           string
	ItbiIncidenciaValor string
}

const (
	outorgadoStartOffSet = "Outorgado"
	outorgadoEndOffSet   = "EscrituraAssinada na Serventia"

	outorganteStartOffSet = "Outorgante"
	outorganteEndOffSet   = "OutorgadoParte"
)

var (
	separators = []string{
		"Parte",
		// "Conjuge",
	}
)

func (e *Extractor) extractPersons(text string, startOffSet, endOffSet string, separators []string) []Person {
	var result []Person

	text, err := cutOffSet(text, startOffSet, endOffSet)
	if err != nil {
		log.Fatal(err)
	}

	persons := strings.Split(text, separators[0])

	for _, p := range persons {
		if p == "" {
			continue
		}
		result = append(result, e.extractPerson(p))
	}

	return result
}

func (e *Extractor) extractPerson(text string) Person {
	tokens := NewPersonTokens2()
	identifierByToken := make(map[Identifier]*Token)

	for _, t := range tokens {
		identifierByToken[t.Identifier] = t

		val, err := extractTokenValue(text, *t)
		if err != nil {
			log.Printf("extract token val err - token: %s - err: %s", IdentifiersNames[t.Identifier], err)
			e.TokensNotFound = append(e.TokensNotFound, t)
			continue
		}

		t.Value = val
		t.IsExtracted = true
		identifierByToken[t.Identifier] = t
	}

	return Person{
		Name:            identifierByToken[Name].Value,
		Nationality:     identifierByToken[Nationality].Value,
		Job:             identifierByToken[Job].Value,
		MaritalStatus:   identifierByToken[MaritialStatus].Value,
		DocNum_CPF_CNPJ: identifierByToken[DocNumCPF_CNPJ].Value,
		DocType:         identifierByToken[DocType].Value,
		Sex:             identifierByToken[Nationality].Value,
		Address: Address{
			Street:       identifierByToken[AddressStreet].Value,
			Num:          identifierByToken[AddressN].Value,
			CityUF:       identifierByToken[AddressCityUF].Value,
			Neighborhood: identifierByToken[AddressNeighborhood].Value,
		},
	}
}

func (e *Extractor) NewExtract(text string) {
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

func New2() *Extractor {
	var tokens []*Token
	tokens = append(tokens, NewScriptureTokens()...)

	return &Extractor{
		tokens:            tokens,
		outorgadoPersons:  nil,
		outorgantePersons: nil,
		TokensNotFound:    nil,
		result2:           make(map[Identifier]*Token),
	}
}

func (e *Extractor) Extract2(text string) {
	for _, token := range e.tokens {
		if token.IsExtracted {
			continue
		}

		e.result2[token.Identifier] = token

		val, err := extractTokenValue(text, *token)
		if err != nil {
			log.Printf("extract token val err - token: %s - err: %s", IdentifiersNames[token.Identifier], err)
			continue
		}

		token.Value = val
		token.IsExtracted = true
		e.result2[token.Identifier] = token
	}

	e.outorgantePersons = append(
		e.outorgantePersons,
		e.extractPersons(text, outorganteStartOffSet, outorganteEndOffSet, separators)...,
	)

	e.outorgadoPersons = append(
		e.outorgadoPersons,
		e.extractPersons(text, outorgadoStartOffSet, outorgadoEndOffSet, separators)...,
	)
}

func (e *Extractor) Result2() Extracted {
	var result Extracted

	for _, t := range e.tokens {
		if !t.IsExtracted {
			log.Printf("token not found - token: '%s'", IdentifiersNames[t.Identifier])
			e.TokensNotFound = append(e.TokensNotFound, t)
		}

		if len(t.Value) >= 55 {
			log.Printf("maybe token value is incorrect - token: '%s'", IdentifiersNames[t.Identifier])
			log.Printf("token value: '%s'", t.Value)
		}
	}

	result.Scripture = Scripture{
		Outorgantes: e.outorgantePersons,
		Outorgados:  e.outorgadoPersons,
		Tablionato: Tablionato{
			Name:   e.result2[TabelionatoName].Value,
			CityUF: e.result2[TabelionatoCityUF].Value,
		},
		TitleAto:            e.result2[TitleAto].Value,
		BookNum:             e.result2[BookNum].Value,
		InitialBookPages:    e.result2[InitialBookPages].Value,
		FinalBookPages:      e.result2[FinalBookPages].Value,
		EscrituraMadeDate:   e.result2[EscrituraMadeDate].Value,
		EscrituraValor:      e.result2[EscrituraValor].Value,
		ItbiValor:           e.result2[ItbiValor].Value,
		ItbiIncidenciaValor: e.result2[ItbiIncidenciaValor].Value,
	}

	return result
}

func NewPersonTokens2() []*Token {
	return []*Token{
		{
			StartKeys:   []string{"Parte :"},
			EndKeys:     []string{"Pessoa:", "Data de Nascimento:"},
			Identifier:  Name,
			IsExtracted: false,
			Value:       defaultValue(Name),
		},
		{
			StartKeys:   []string{"Profissão:"},
			EndKeys:     []string{" - Nacionalidade:"},
			Identifier:  Job,
			IsExtracted: false,
			Value:       defaultValue(Job),
		},
		{
			StartKeys:   []string{"Nacionalidade: "},
			EndKeys:     []string{"- Sexo:", " -"},
			Identifier:  Nationality,
			IsExtracted: false,
			Value:       defaultValue(Nationality),
		},
		{
			StartKeys:   []string{"Estado Civil: "},
			EndKeys:     []string{" -"},
			Identifier:  MaritialStatus,
			IsExtracted: false,
			Value:       defaultValue(MaritialStatus),
		},
		{
			StartKeys:   []string{"Sexo:"},
			EndKeys:     []string{"DocumentosDoc."},
			Identifier:  Sex,
			IsExtracted: false,
			Value:       defaultValue(Sex),
		},
		{
			StartKeys:   []string{"Doc. Nº:"},
			EndKeys:     []string{"/", "Doc. Tipo:", "EndereçosLogradouro:"},
			StartOffset: "Outorgante",
			Identifier:  DocNumCPF_CNPJ,
			IsExtracted: false,
			Value:       defaultValue(DocNumCPF_CNPJ),
		},
		{
			StartKeys:   []string{"Doc. Tipo:", "DocumentosDoc. Tipo: "},
			EndKeys:     []string{"Doc."},
			Identifier:  DocType,
			IsExtracted: false,
			Value:       defaultValue(DocType),
		},
		{
			StartKeys:   []string{"EndereçosLogradouro:"},
			EndKeys:     []string{"Número:"},
			Identifier:  AddressStreet,
			IsExtracted: false,
			Value:       defaultValue(AddressStreet),
		},
		{
			StartKeys:   []string{"Número:"},
			EndKeys:     []string{"Bairro:"},
			Identifier:  AddressN,
			IsExtracted: false,
			Value:       defaultValue(AddressN),
		},
		{
			StartKeys:   []string{"Bairro:"},
			EndKeys:     []string{"Complemento", "Cidade/UF:", "CEP:", "Cidade:", ","},
			Identifier:  AddressNeighborhood,
			IsExtracted: false,
			Value:       defaultValue(AddressNeighborhood),
		},
		{
			StartKeys:   []string{"Cidade/UF: "},
			EndKeys:     []string{"CEP:"},
			Identifier:  AddressCityUF,
			IsExtracted: false,
			Value:       defaultValue(AddressCityUF),
		},
	}
}

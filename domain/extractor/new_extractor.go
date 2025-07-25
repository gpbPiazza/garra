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
	// InFavorToWho is a string where contains the list of all persons
	// favority in this scripture, there we can extract informations about the
	// relation between the persons.
	InFavorToWho string
}

type Tablionato struct {
	Name   string
	CityUF string
}

type Party struct {
	Persons       []Person
	Beneficiaries string // List of names who benefit from this party's actions
}

type Scripture struct {
	Outorgantes Party
	Outorgados  Party
	Tablionato  Tablionato

	BookNum          string
	InitialBookPages string
	FinalBookPages   string

	TitleAto            string
	EscrituraMadeDate   string
	EscrituraValor      string
	ItbiValor           string
	ItbiIncidenciaValor string
}

const (
	outorgadoStartOffSet = "Outorgado"
	outorgadoEndOffSet   = "EscrituraAssinada na Serventia"

	outorganteStartOffSet = "Outorgante"
	outorganteEndOffSet   = "Outorgado"
)

var (
	separators = []string{
		"Parte",
		// "Conjuge",
	}
)

func (e *Extractor) extractPersons(text string, startOffSet, endOffSet string, separators []string) []Person {
	var result []Person

	personTxt, err := cutOffSet(text, startOffSet, endOffSet)
	if err != nil {
		log.Fatal(err)
	}

	persons := strings.SplitAfter(personTxt, separators[0])

	for _, p := range persons {
		if p == "" || strings.Contains(p, "ParcialParte") || strings.HasPrefix(p, "Parte") {
			continue
		}
		result = append(result, e.extractPerson(p))
	}

	return result
}

func (e *Extractor) extractPerson(personOffSetTxt string) Person {
	tokens := NewPersonTokens2()
	identifierByToken := make(map[Identifier]*Token)

	for _, t := range tokens {
		identifierByToken[t.Identifier] = t

		val, err := extractTokenValue(personOffSetTxt, *t)
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
		Name:            identifierByToken[NameID].Value,
		Nationality:     identifierByToken[NationalityID].Value,
		Job:             identifierByToken[JobID].Value,
		MaritalStatus:   identifierByToken[MaritialStatusID].Value,
		DocNum_CPF_CNPJ: identifierByToken[DocNumCPF_CNPJID].Value,
		DocType:         identifierByToken[DocTypeID].Value,
		Sex:             identifierByToken[SexID].Value,
		Address: Address{
			Street:       identifierByToken[AddressStreetID].Value,
			Num:          identifierByToken[AddressNID].Value,
			CityUF:       identifierByToken[AddressCityUFID].Value,
			Neighborhood: identifierByToken[AddressNeighborhoodID].Value,
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

	result.TokensNotFound = e.TokensNotFound
	result.Scripture = Scripture{
		Outorgantes: Party{Persons: e.outorgantePersons, Beneficiaries: e.result2[WhoDoesID].Value},
		Outorgados:  Party{Persons: e.outorgadoPersons, Beneficiaries: e.result2[InFavorToWhoID].Value},
		Tablionato: Tablionato{
			Name:   e.result2[TabelionatoNameID].Value,
			CityUF: e.result2[TabelionatoCityUFID].Value,
		},
		TitleAto:            e.result2[TitleAtoID].Value,
		BookNum:             e.result2[BookNumID].Value,
		InitialBookPages:    e.result2[InitialBookPagesID].Value,
		FinalBookPages:      e.result2[FinalBookPagesID].Value,
		EscrituraMadeDate:   e.result2[EscrituraMadeDateID].Value,
		EscrituraValor:      e.result2[EscrituraValorID].Value,
		ItbiValor:           e.result2[ItbiValorID].Value,
		ItbiIncidenciaValor: e.result2[ItbiIncidenciaValorID].Value,
	}

	return result
}

func NewPersonTokens2() []*Token {
	return []*Token{
		{
			StartKeys:   []string{"Parte :", ":"},
			EndKeys:     []string{"Pessoa:", "Data de Nascimento:"},
			Identifier:  NameID,
			IsExtracted: false,
			Value:       defaultValue(NameID),
		},
		{
			StartKeys:   []string{"Profissão:"},
			EndKeys:     []string{" - Nacionalidade:"},
			Identifier:  JobID,
			IsExtracted: false,
			Value:       defaultValue(JobID),
		},
		{
			StartKeys:   []string{"Nacionalidade: "},
			EndKeys:     []string{"- Sexo:", " -"},
			Identifier:  NationalityID,
			IsExtracted: false,
			Value:       defaultValue(NationalityID),
		},
		{
			StartKeys:   []string{"Estado Civil: "},
			EndKeys:     []string{" -", "- Profissão"},
			Identifier:  MaritialStatusID,
			IsExtracted: false,
			Value:       defaultValue(MaritialStatusID),
		},
		{
			StartKeys:   []string{"Sexo:"},
			EndKeys:     []string{"DocumentosDoc."},
			Identifier:  SexID,
			IsExtracted: false,
			Value:       defaultValue(SexID),
		},
		{
			StartKeys:   []string{"Doc. Nº:"},
			EndKeys:     []string{"/", "Doc. Tipo:", "EndereçosLogradouro:"},
			Identifier:  DocNumCPF_CNPJID,
			IsExtracted: false,
			Value:       defaultValue(DocNumCPF_CNPJID),
		},
		{
			StartKeys:   []string{"Doc. Tipo:", "DocumentosDoc. Tipo: "},
			EndKeys:     []string{"Doc."},
			Identifier:  DocTypeID,
			IsExtracted: false,
			Value:       defaultValue(DocTypeID),
		},
		{
			StartKeys:   []string{"EndereçosLogradouro:"},
			EndKeys:     []string{"Número:"},
			Identifier:  AddressStreetID,
			IsExtracted: false,
			Value:       defaultValue(AddressStreetID),
		},
		{
			StartKeys:   []string{"Número:"},
			EndKeys:     []string{"Bairro:"},
			Identifier:  AddressNID,
			IsExtracted: false,
			Value:       defaultValue(AddressNID),
		},
		{
			StartKeys:   []string{"Bairro:"},
			EndKeys:     []string{"Complemento", "Cidade/UF:", "CEP:", "Cidade:", ","},
			Identifier:  AddressNeighborhoodID,
			IsExtracted: false,
			Value:       defaultValue(AddressNeighborhoodID),
		},
		{
			StartKeys:   []string{"Cidade/UF: "},
			EndKeys:     []string{"CEP:"},
			Identifier:  AddressCityUFID,
			IsExtracted: false,
			Value:       defaultValue(AddressCityUFID),
		},
	}
}

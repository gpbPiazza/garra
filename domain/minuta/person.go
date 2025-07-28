package minuta

import (
	"fmt"
	"log"
	"strings"

	"github.com/gpbPiazza/garra/domain/extractor"
)

type AddressParams struct {
	Rua          string
	Num          string
	CityUF       string
	Neighborhood string
}

type PersonParams struct {
	extractor.Person
	IsOverqualified bool
}

func IsJuridicPerson(docType string) bool {
	return strings.EqualFold("CNPJ", docType)
}

func IsFisicalPerson(docType string) bool {
	return strings.EqualFold("CPF", docType)
}

func minutaPerson(person PersonParams) (string, error) {
	if IsJuridicPerson(person.DocType) {
		return juridicPerson(person)
	}

	if IsFisicalPerson(person.DocType) {
		return fisicalPerson(person)
	}

	return "", fmt.Errorf("docType not mapped - type %s", person.DocType)
}

func juridicPerson(person PersonParams) (string, error) {
	name := formatName(person.Name)

	if person.IsOverqualified {
		return fmt.Sprintf("<strong>%s.</strong>, supraqualificada.", name), nil
	}

	doc, err := formatCNPJDoc(person.DocNum_CPF_CNPJ)
	if err != nil {
		log.Printf("err to format person doc err: %s", err)
		return "", err
	}

	cityUF, err := formatCityUF(person.Address.CityUF)
	if err != nil {
		log.Printf("err on formatCityUF err: %s", err)
		return "", err
	}

	neighborhood, err := formatNeighborhood(person.Address.Neighborhood)
	if err != nil {
		log.Printf("err on formatNeighborhood err: %s", err)
		return "", err
	}

	street := formatStreet(person.Address.Street)

	return fmt.Sprintf(
		"<strong>%s.</strong>, CNPJ nº %s, com sede na rua %s, nº %s, Bairro %s, %s.",
		name,
		doc,
		street,
		person.Address.Num,
		neighborhood,
		cityUF,
	), nil
}

func fisicalPerson(person PersonParams) (string, error) {
	name := formatName(person.Name)

	if person.IsOverqualified {
		return fmt.Sprintf("<strong>%s</strong>, supraqualificada.", name), nil
	}

	doc, err := formatCPFDoc(person.DocNum_CPF_CNPJ)
	if err != nil {
		log.Printf("err to format person doc err: %s", err)
		return "", err
	}

	nationality, err := formatNationality(person.Nationality, person.Sex)
	if err != nil {
		log.Printf("err on formatNationality err: %s", err)
		return "", err
	}

	maritialStatus, err := formatMaritialStatus(person.MaritalStatus, person.Sex)
	if err != nil {
		log.Printf("err on formatMaritialStatus err: %s", err)
		return "", err
	}

	address, err := minutaAdress(person.Address)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"<strong>%s</strong>, %s, %s, %s, CPF nº %s, residente e domiciliado na %s",
		name,
		nationality,
		maritialStatus,
		formatJob(person.Job),
		doc,
		address,
	), nil
}

func minutaMarried(party extractor.Party, _ bool) (string, error) {
	man, woman := splitMarried(party.Persons)

	manMinuta, err := minutaMarriedPerson(man)
	if err != nil {
		return "", err
	}

	womanMinuta, err := minutaMarriedPerson(woman)
	if err != nil {
		return "", err
	}

	addressMinuta, err := minutaAdress(man.Address)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s, e sua esposa %s, casados pelo regime da Comunhão Parcial de Bens, na vigência da Lei n° 6.515/77, residentes e domiciliados na %s",
		manMinuta,
		womanMinuta,
		addressMinuta,
	), nil
}

func minutaMarriedPerson(person extractor.Person) (string, error) {
	name := formatName(person.Name)

	doc, err := formatCPFDoc(person.DocNum_CPF_CNPJ)
	if err != nil {
		log.Printf("err to format person doc err: %s", err)
		return "", err
	}

	// TODO paramos aqui -> person.Sex da mulher está vindo como masculino
	// Estamos corrigindo o teste unitário de muitos pra muitos
	// Notei também que eu estou marretando o tipo de regime do casamento, preciso ajustar para
	// extrair do DOC qual o regime.
	nationality, err := formatNationality(person.Nationality, person.Sex)
	if err != nil {
		log.Printf("err on formatNationality err: %s", err)
		return "", err
	}

	return fmt.Sprintf(
		"<strong>%s</strong>, %s, %s, CPF nº %s",
		name,
		nationality,
		formatJob(person.Job),
		doc,
	), nil
}

func splitMarried(persons []extractor.Person) (man, woman extractor.Person) {
	if len(persons) != 2 {
		panic("split married tried to split more than 2 two persons slice")
	}

	if Lower(persons[0].Sex) == "masculino" {
		man = persons[0]
		woman = persons[1]

		return
	}

	man = persons[1]
	woman = persons[0]
	return
}

func minutaAdress(address extractor.Address) (string, error) {
	cityUF, err := formatCityUF(address.CityUF)
	if err != nil {
		log.Printf("err on formatCityUF err: %s", err)
		return "", err
	}

	neighborhood, err := formatNeighborhood(address.Neighborhood)
	if err != nil {
		log.Printf("err on formatNeighborhood err: %s", err)
		return "", err
	}

	street := formatStreet(address.Street)

	return fmt.Sprintf("%s, nº %s, Bairro %s, %s.", street, address.Num, neighborhood, cityUF), nil
}

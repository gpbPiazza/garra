package minuta

import (
	"fmt"
	"log"
	"strings"
)

type AddressParams struct {
	Rua          string
	Num          string
	CityUF       string
	Neighborhood string
}

type PersonParams struct {
	Name            string
	Nationality     string
	Job             string
	MaritalStatus   string
	DocNum_CPF_CNPJ string
	DocType         string
	Address         AddressParams
	IsOverqualified bool
	Sex             string
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

	street := formatStreet(person.Address.Rua)

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

	nationality, err := formatNationality(person.Nationality)
	if err != nil {
		log.Printf("err on formatNationality err: %s", err)
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

	street := formatStreet(person.Address.Rua)

	maritialStatus, err := formatMaritialStatus(person.MaritalStatus, person.Sex)
	if err != nil {
		log.Printf("err on formatMaritialStatus err: %s", err)
		return "", err
	}

	return fmt.Sprintf(
		"<strong>%s</strong>, %s, %s, %s, CPF nº %s, residente e domiciliado na %s, nº %s, Bairro %s, %s.",
		name,
		nationality,
		maritialStatus,
		formatJob(person.Job),
		doc,
		street,
		person.Address.Num,
		neighborhood,
		cityUF,
	), nil
}

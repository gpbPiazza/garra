package minuta

import (
	"errors"
	"fmt"
	"log"
	"regexp"
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
	MaritalStatus   string
	DocNum_CPF_CNPJ string
	DocType         string
	Address         AddressParams
	IsOverqualified bool
}

func minutaPerson(person PersonParams) (string, error) {
	switch strings.ToUpper(person.DocType) {
	case "CNPJ":
		return juridicPerson(person)
	case "CPF":
		return fisicalPerson(person)
	default:
		return "", fmt.Errorf("docType not mapped - type %s", person.DocType)
	}
}

func juridicPerson(person PersonParams) (string, error) {
	if person.IsOverqualified {
		return fmt.Sprintf("%s., supraqualificada.", person.Name), nil
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

	return fmt.Sprintf(
		"%s., CNPJ nº %s, com sede na rua %s, nº %s, Bairro %s, %s.",
		person.Name,
		doc,
		person.Address.Rua,
		person.Address.Num,
		neighborhood,
		cityUF,
	), nil
}

func fisicalPerson(person PersonParams) (string, error) {
	if person.IsOverqualified {
		return fmt.Sprintf("%s, supraqualificada.", person.Name), nil
	}

	doc, err := formatCPFDoc(person.DocNum_CPF_CNPJ, person.DocType)
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

	return fmt.Sprintf(
		"%s, %s, %s, CPF nº %s, residente e domiciliado na %s, nº %s, Bairro %s, %s.",
		person.Name,
		nationality,
		person.MaritalStatus,
		doc,
		person.Address.Rua,
		person.Address.Num,
		neighborhood,
		cityUF,
	), nil
}

func formatCityUF(cityUF string) (string, error) {
	cityUFSplt := strings.Split(cityUF, "/")
	if len(cityUFSplt) != 2 {
		return "", errors.New("cityUF can not be splited by /")
	}

	city := strings.TrimSpace(cityUFSplt[0])
	UF := strings.TrimSpace(cityUFSplt[1])

	return fmt.Sprintf("%s/%s", city, UF), nil
}

func formatNationality(nationality string) (string, error) {
	var fNationality string

	switch strings.ToLower(nationality) {
	case "brasil":
		fNationality = "brasileiro"
	default:
		return fNationality, fmt.Errorf("nationality not mapped - got %s", nationality)
	}

	return fNationality, nil
}

func formatCNPJDoc(docValue string) (string, error) {
	var formattedDoc string
	docValue = strings.TrimSpace(docValue)

	if len(docValue) == 16 {
		// we have a problem in the extractor thata alwys get two digits more some times
		// this is quick fix for now
		docValue = docValue[:len(docValue)-2]
	}

	if len(docValue) != 14 {
		return formattedDoc, errors.New("malformed CNPJ value len diff than 14")
	}

	cnpjRegex := regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`)
	formattedDoc = cnpjRegex.ReplaceAllString(docValue, "$1.$2.$3/$4-$5")

	return formattedDoc, nil
}

func formatCPFDoc(docValue string, docType string) (string, error) {
	var formattedDoc string
	docValue = strings.TrimSpace(docValue)

	if len(docValue) != 11 {
		return formattedDoc, errors.New("malformed CPF value len diff than 11")
	}
	cpfRegex := regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
	formattedDoc = cpfRegex.ReplaceAllString(docValue, "$1.$2.$3-$4")

	return formattedDoc, nil
}

func formatNeighborhood(neighborhood string) (string, error) {
	if strings.Contains(neighborhood, "/") {
		return removeDateFromStr(neighborhood), nil
	}

	return neighborhood, nil
}

// Ex: "POÇO FUNDO24/04/2025" OUTPUT: "POÇO FUNDO"
func removeDateFromStr(str string) string {
	nSplited := strings.Split(str, "/")
	str = nSplited[0]
	str = str[:len(str)-2]
	return str
}

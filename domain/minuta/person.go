package minuta

import (
	"errors"
	"fmt"
	"log"
	"regexp"
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
	Name            string
	Nationality     string
	MaritalStatus   string
	DocNum_CPF_CNPJ string
	DocType         string
	Address         AddressParams
	IsOverqualified bool
	Sex             string
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
		"<strong>%s</strong>, %s, %s, CPF nº %s, residente e domiciliado na %s, nº %s, Bairro %s, %s.",
		name,
		nationality,
		maritialStatus,
		// "profissão do cara",
		doc,
		street,
		person.Address.Num,
		neighborhood,
		cityUF,
	), nil
}

func formatName(name string) string {
	if notFound(name) {
		return name
	}

	return strings.TrimSpace(strings.ToUpper(name))
}

func formatMaritialStatus(maritialStatus string, personSex string) (string, error) {
	if notFound(maritialStatus) {
		return maritialStatus, nil
	}

	if maritialStatus == "" {
		return "", errors.New("error - maritialStatus is required")
	}
	if personSex == "" {
		return "", errors.New("error - personSex is required")
	}

	maritialStatus = strings.ToLower(maritialStatus)
	if strings.Contains(maritialStatus, "separado") ||
		strings.Contains(maritialStatus, "divorciado") {
		switch strings.ToLower(personSex) {
		case "masculino":
			return "divorciado", nil
		default:
			return "divorciada", nil
		}
	}

	if strings.Contains(maritialStatus, "solteiro") {
		switch strings.ToLower(personSex) {
		case "masculino":
			return "solteiro", nil
		default:
			return "solteira", nil
		}
	}

	if strings.Contains(maritialStatus, "casado") {
		switch strings.ToLower(personSex) {
		case "masculino":
			return "casado", nil
		default:
			return "casada", nil
		}
	}

	return "", errors.New("UNKNOW cases to formatMaritialStatus")
}

func formatCityUF(cityUF string) (string, error) {
	if notFound(cityUF) {
		return cityUF, nil
	}

	cityUFSplt := strings.Split(cityUF, "/")
	if len(cityUFSplt) != 2 {
		return "", errors.New("cityUF can not be splited by /")
	}

	city := strings.TrimSpace(cityUFSplt[0])
	UF := strings.TrimSpace(cityUFSplt[1])

	return fmt.Sprintf("%s/%s", city, UF), nil
}

func formatNationality(nationality string) (string, error) {
	if notFound(nationality) {
		return nationality, nil
	}

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
	if notFound(docValue) {
		return docValue, nil
	}

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

func formatCPFDoc(docValue string) (string, error) {
	if notFound(docValue) {
		return docValue, nil
	}

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
	if notFound(neighborhood) {
		return neighborhood, nil
	}

	if strings.Contains(neighborhood, "/") {
		return removeDateFromStr(neighborhood), nil
	}

	return capitalizeEachWord(neighborhood), nil
}

// Ex: "POÇO FUNDO24/04/2025" OUTPUT: "POÇO FUNDO"
func removeDateFromStr(str string) string {
	if notFound(str) {
		return str
	}

	nSplited := strings.Split(str, "/")
	str = nSplited[0]
	str = str[:len(str)-2]
	return str
}

func notFound(val string) bool {
	return strings.Contains(val, extractor.NotFoundDefaultSuffix)
}

func formatStreet(street string) string {
	street = capitalizeEachWord(street)

	words := strings.Split(street, " ")

	for i, word := range words {
		if strings.EqualFold(word, "De") {
			words[i] = Lower(word)
			continue
		}

		if len(word) <= 2 {
			words[i] = Title(word)
		}
	}

	return strings.Join(words, " ")
}

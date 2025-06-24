package minuta

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gpbPiazza/garra/domain/extractor"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func capitalizeEachWord(sentence string) string {
	if notFound(sentence) {
		return sentence
	}

	words := strings.Split(sentence, " ")

	for i, word := range words {
		if len(word) == 0 {
			continue
		}

		word = Lower(word)

		isLetter := len(word) == 1
		isPreposition := len(word) == 2
		if isLetter || isPreposition {
			words[i] = word
			continue
		}

		words[i] = Title(word)
	}

	return strings.Join(words, " ")
}

func formatDate(dateStr string) (string, error) {
	if notFound(dateStr) {
		return dateStr, nil
	}

	if len(dateStr) == 24 { // This case formats the EscrituraMadeDate whenever the end key ends with the page
		dateStr = dateStr[:len(dateStr)-10]
	}

	// 26 / 03 / 202524/04/2025

	dateSplit := strings.Split(dateStr, "/")
	var date []string

	if len(dateSplit) != 3 {
		return "", fmt.Errorf("can not split by / to formatDate dateStr: %s", dateStr)
	}

	for _, d := range dateSplit {
		date = append(date, strings.TrimSpace(d))
	}

	return strings.Join(date, "/"), nil
}

// Código feito por IA cuidado bixo
func formatValue(val string) (string, error) {
	if notFound(val) {
		return val, nil
	}

	// Check if input is empty
	if val == "" {
		return "", fmt.Errorf("empty value provided")
	}

	// Step 1: Extract the numeric part (R$ xxx.xxx,xx)
	numericPart := ""
	idx := 0
	for idx < len(val) && (val[idx] != '(' && idx < len(val)-1) {
		numericPart += string(val[idx])
		idx++
	}
	numericPart = strings.TrimSpace(numericPart)

	// Normalize the numeric part format (ensure space after R$)
	if strings.HasPrefix(numericPart, "R$") && !strings.HasPrefix(numericPart, "R$ ") {
		numericPart = "R$ " + numericPart[2:]
	}

	// Step 2: Extract the text description part between parentheses
	textStart := strings.Index(val, "(")
	if textStart == -1 {
		return "", fmt.Errorf("missing opening parenthesis in value: %s", val)
	}

	textEnd := strings.LastIndex(val, ")")
	var textPart string

	// Handle case where closing parenthesis is missing
	if textEnd == -1 {
		// Extract text from opening parenthesis to the end
		rawText := val[textStart:]

		// Look for date pattern and remove it
		dateParts := strings.Split(rawText, "mil")
		if len(dateParts) > 1 {
			// Keep only the first part that has the actual amount in words
			textPart = dateParts[0] + "mil reais)"
		} else {
			// If no "mil" keyword, just append proper ending
			textPart = rawText + " reais)"
		}
	} else {
		// Normal case with proper closing parenthesis
		textPart = val[textStart : textEnd+1]
	}

	// Fix common typos in text part
	textPart = strings.ReplaceAll(textPart, "edez", "e dez")

	// Check if "reais" is missing at the end
	if !strings.Contains(strings.ToLower(textPart), "reais") {
		// Remove closing parenthesis if exists
		textPart = strings.TrimSuffix(textPart, ")")
		textPart += " reais)"
	}

	// Final formatting
	return fmt.Sprintf("%s %s", numericPart, textPart), nil
}

func Upper(str string) string {
	return cases.Upper(language.BrazilianPortuguese, cases.Compact).String(str)
}

func Lower(str string) string {
	return cases.Lower(language.BrazilianPortuguese, cases.Compact).String(str)
}

func Title(str string) string {
	return cases.Title(language.BrazilianPortuguese, cases.Compact).String(str)
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
		neighborhood = removeDateFromStr(neighborhood)
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

func formatJob(job string) string {
	if notFound(job) {
		return job
	}

	return Lower(job)
}

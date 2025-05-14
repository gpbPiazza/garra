package minuta

import (
	"fmt"
	"strings"
)

const minutaTemplate = `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>{{TRANSMITENTE}}</span>
<br/>
<u>ADQUIRENTE(S)</u>:><span>{{ADQUIRENTE}}</span>
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de {{TITLE_ATO}}, lavrada pelo {{TABELIONATO_NAME}}` +
	` de {{TABELIONATO_CITY_STATE}}, Livro {{BOOK_NUM}}, Folhas {{INITIAL_BOOK_PAGES}}/{{FINAL_BOOK_PAGES}}V, em {{BOOK_DATE}}. 
<br/><u>VALOR</u>: R$ {{ESCRITURA_VALOR}}.
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ {{ITBI_VALOR}}, com incidência ` +
	`sobre R$ {{ITBI_INCIDENCIA_VALOR}}, devidamente quitado.` +
	` No ato da lavratura da Escritura Pública, foram apresentadas as certidões ` +
	`previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

// Template com contestação
// `sobre R$ {{ITBI_INCIDENCIA_VALOR}}, devidamente quitado. Nos termos do artigo 320 do CNCGFE/SC, ` +
// `o imóvel da presente matrícula, teve como valor atribuído de mercado, no
// ''quantum'' de R$ XXXXX. No ato da lavratura da Escritura Pública, foram apresentadas as certidões ` +

type MinutaParams struct {
	Transmitente        PersonParams
	Adquirente          PersonParams
	TitleAto            string
	TabelionatoName     string
	TabelionatoCityUF   string
	BookNum             string
	InitialBookPages    string
	FinalBookPages      string
	EscrituraMadeDate   string
	EscrituraValor      string
	ItbiValor           string
	ItbiIncidenciaValor string
}

func Minuta(params MinutaParams) (string, error) {
	transmitante, err := minutaPerson(params.Transmitente)
	if err != nil {
		return "", err
	}

	adquidirente, err := minutaPerson(params.Adquirente)
	if err != nil {
		return "", err
	}

	tabelionatoCityUF, err := formatCityUF(params.TabelionatoCityUF)
	if err != nil {
		return "", err
	}

	escrituraMadeDate, err := formatDate(params.EscrituraMadeDate)
	if err != nil {
		return "", err
	}

	value, err := formatValue(params.EscrituraValor)
	if err != nil {
		return "", err
	}

	replacer := strings.NewReplacer(
		Transmitente.String(), transmitante,
		Adquirente.String(), adquidirente,
		TitleAto.String(), capitalizeEachWord(params.TitleAto),
		TabelionatoName.String(), capitalizeEachWord(params.TabelionatoName),
		TabelionatoCityUF.String(), tabelionatoCityUF,
		BookNum.String(), params.BookNum,
		InitialBookPages.String(), params.InitialBookPages,
		FinalBookPages.String(), params.FinalBookPages,
		EscrituraMadeDate.String(), escrituraMadeDate,
		EscrituraValor.String(), value,
		ItbiValor.String(), params.ItbiValor,
		ItbiIncidenciaValor.String(), params.ItbiIncidenciaValor,
	)

	return replacer.Replace(minutaTemplate), nil
}

func capitalizeEachWord(sentence string) string {
	words := strings.Split(sentence, " ")

	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		word = strings.ToLower(word)

		isLetter := len(word) == 1
		isPreposition := len(word) == 2
		if isLetter || isPreposition {
			words[i] = word
			continue
		}

		firstLetter := strings.ToUpper(string(word[0]))
		words[i] = firstLetter + word[1:]
	}

	return strings.Join(words, " ")
}

func formatDate(dateStr string) (string, error) {
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

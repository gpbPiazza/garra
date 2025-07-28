package minuta

import (
	"errors"
	"strings"

	"github.com/gpbPiazza/garra/domain/extractor"
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

func Minuta(params extractor.Scripture, IsTransOverqualified, IsAdquiOverqualified bool) (string, error) {
	transmitante, _ := minutaPersons(params.Outorgantes, IsTransOverqualified)

	adquidirente, _ := minutaPersons(params.Outorgados, IsAdquiOverqualified)

	tabelionatoCityUF, _ := formatCityUF(params.Tablionato.CityUF)

	escrituraMadeDate, _ := formatDate(params.EscrituraMadeDate)

	value, _ := formatValue(params.EscrituraValor)

	replacer := strings.NewReplacer(
		TransmitenteRK.String(), transmitante,
		AdquirenteRK.String(), adquidirente,
		TitleAtoRK.String(), capitalizeEachWord(params.TitleAto),
		TabelionatoNameRK.String(), capitalizeEachWord(params.Tablionato.Name),
		TabelionatoCityUFRK.String(), tabelionatoCityUF,
		BookNumRK.String(), params.BookNum,
		InitialBookPagesRK.String(), params.InitialBookPages,
		FinalBookPagesRK.String(), params.FinalBookPages,
		EscrituraMadeDateRK.String(), escrituraMadeDate,
		EscrituraValorRK.String(), value,
		ItbiValorRK.String(), params.ItbiValor,
		ItbiIncidenciaValorRK.String(), params.ItbiIncidenciaValor,
	)

	return replacer.Replace(minutaTemplate), nil
}

func minutaPersons(party extractor.Party, isOverQualified bool) (string, error) {
	if len(party.Persons) == 1 {
		return minutaPerson(PersonParams{IsOverqualified: isOverQualified, Person: party.Persons[0]})
	}

	isMarried, _ := isTwoPartiesMarried(party)
	if isMarried {
		return minutaMarried(party, isOverQualified)
	}

	return "", errors.New("unknow minutaPersons case")
}

func isTwoPartiesMarried(party extractor.Party) (bool, error) {
	if notFound(party.Beneficiaries) {
		return false, nil
	}

	if party.Beneficiaries == "" {
		return false, errors.New("error - party.Beneficiaries is required")
	}

	benef := Lower(party.Beneficiaries)

	if !strings.Contains(benef, "esposa") && !strings.Contains(benef, "marido") {
		return false, nil
	}

	benefs := strings.Split(benef, "e sua esposa")

	isTwoPersonsMarried := len(benefs) == 2 &&
		len(party.Persons) == 2

	return isTwoPersonsMarried, nil
}

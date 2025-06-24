package minuta

import (
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
	transmitante, _ := minutaPerson(params.Transmitente)

	adquidirente, _ := minutaPerson(params.Adquirente)

	tabelionatoCityUF, _ := formatCityUF(params.TabelionatoCityUF)

	escrituraMadeDate, _ := formatDate(params.EscrituraMadeDate)

	value, _ := formatValue(params.EscrituraValor)

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

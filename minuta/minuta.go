package minuta

import (
	"strings"
)

const minutaTemplate = `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>{{TRANSMITENTE}}</span>.
<br/>
<u>ADQUIRENTE(S)</u>:><span>{{ADQUIRENTE}}</span>.
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de {{TITLE_ATO}}, lavrada pelo {{TABELIONATO_NAME}} de {{TABELIONATO_CITY_STATE}}, Livro {{BOOK_NUM}}, Folhas {{BOOK_PAGES}}, em {{BOOK_DATE}}. 
<br/><u>VALOR</u>: R$ {{ESCRITURA_VALOR}} {{ESCRITURA_VALOR_EXTENSO}}.
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ {{ITBI_VALOR}}, com incidência sobre R$ {{ITBI_INCIDENCIA_VALOR}}, devidamente quitado. Nos termos do artigo 320 do CNCGFE/SC, o imóvel da presente matrícula, teve como valor atribuído de mercado, no
''quantum'' de R$ XXXXX. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com
as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

type MinutaParams struct {
	Transmitente          string
	Adquirente            string
	TitleAto              string
	TabelionatoName       string
	TabelionatoCityState  string
	BookNum               string
	BookPages             string
	EscrituraMadeDate     string
	EscrituraValor        string
	EscrituraValorExtenso string
	ItbiValor             string
	ItbiIncidenciaValor   string
}

func Minuta(params MinutaParams) string {
	replacer := strings.NewReplacer(
		Transmitente.String(), params.Transmitente,
		Adquirente.String(), params.Adquirente,
		TitleAto.String(), params.TitleAto,
		TabelionatoName.String(), params.TabelionatoName,
		TabelionatoCityState.String(), params.TabelionatoCityState,
		BookNum.String(), params.BookNum,
		EscrituraMadeDate.String(), params.EscrituraMadeDate,
		EscrituraValor.String(), params.EscrituraValor,
		ItbiValor.String(), params.ItbiValor,
		ItbiIncidenciaValor.String(), params.ItbiIncidenciaValor,
	)

	return replacer.Replace(minutaTemplate)
}

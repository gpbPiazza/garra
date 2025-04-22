package minuta

import (
	"fmt"
	"strings"
)

const minutaTemplate = `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>{{TRANSMITENTE}}</span>.
<br/>
<u>ADQUIRENTE(S)</u>:><span>{{ADQUIRENTE}}</span>.
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de {{TITLE_ATO}}, lavrada pelo {{TABELIONATO_NAME}}` +
	` de {{TABELIONATO_CITY_STATE}}, Livro {{BOOK_NUM}}, Folhas {{INITIAL_BOOK_PAGES}}/{{FINAL_BOOK_PAGES}}V, em {{BOOK_DATE}}. 
<br/><u>VALOR</u>: R$ {{ESCRITURA_VALOR}}.
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ {{ITBI_VALOR}}, com incidência ` +
	`sobre R$ {{ITBI_INCIDENCIA_VALOR}}, devidamente quitado. Nos termos do artigo 320 do CNCGFE/SC, ` +
	`o imóvel da presente matrícula, teve como valor atribuído de mercado, no
''quantum'' de R$ XXXXX. No ato da lavratura da Escritura Pública, foram apresentadas as certidões ` +
	`previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

type AddressParams struct {
	Rua          string
	Num          string
	CityUF       string
	Neighborhood string
}

type PersonParams struct {
	Name            string
	Job             string
	Nationality     string
	MaritalStatus   string
	CPF_CNPJ        string
	Address         AddressParams
	IsOverqualified bool
}

type MinutaParams struct {
	Transmitente         PersonParams
	Adquirente           PersonParams
	TitleAto             string
	TabelionatoName      string
	TabelionatoCityState string
	BookNum              string
	InitialBookPages     string
	FinalBookPages       string
	EscrituraMadeDate    string
	EscrituraValor       string
	ItbiValor            string
	ItbiIncidenciaValor  string
}

func minutaPerson(person PersonParams) string {
	if person.IsOverqualified {
		return fmt.Sprintf("%s, supraqualificada.", person.Name)
	}

	// TODO implement: coditions when person is CPNJ.
	// TODO check if residente a domiciliado is just a case ou default.
	return fmt.Sprintf(
		"%s, %s, %s, CPF nº %s, residente e domiciliado na %s, nº %s, %s, %s.",
		person.Name,
		person.Nationality,
		person.MaritalStatus,
		person.CPF_CNPJ,
		person.Address.Rua,
		person.Address.Num,
		person.Address.Neighborhood,
		person.Address.CityUF,
	)
}

func Minuta(params MinutaParams) string {
	replacer := strings.NewReplacer(
		Transmitente.String(), minutaPerson(params.Transmitente),
		Adquirente.String(), minutaPerson(params.Adquirente),
		TitleAto.String(), params.TitleAto,
		TabelionatoName.String(), params.TabelionatoName,
		TabelionatoCityState.String(), params.TabelionatoCityState,
		BookNum.String(), params.BookNum,
		InitialBookPages.String(), params.InitialBookPages,
		FinalBookPages.String(), params.FinalBookPages,
		EscrituraMadeDate.String(), params.EscrituraMadeDate,
		EscrituraValor.String(), params.EscrituraValor,
		ItbiValor.String(), params.ItbiValor,
		ItbiIncidenciaValor.String(), params.ItbiIncidenciaValor,
	)

	return replacer.Replace(minutaTemplate)
}

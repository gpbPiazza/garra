package minuta

type ReplaceKey string

func (r ReplaceKey) String() string {
	return string(r)
}

var (
	// CABECALHO KEYS
	TypeAto       ReplaceKey = "{{TYPE_ATO}}"
	NumAto        ReplaceKey = "{{NUM_ATO}}"
	TitleAto      ReplaceKey = "{{TITLE_ATO}}"
	Matricula     ReplaceKey = "{{MATRICULA}}"
	DataRegistro  ReplaceKey = "{{DATA_REGISTRO}}"
	Protocolo     ReplaceKey = "{{PROTOCOLO}}"
	DataProtocolo ReplaceKey = "{{DATA_PROTOCOLO}}"

	// MINUTA KEYS
	Transmitente          ReplaceKey = "{{TRANSMITENTE}}"
	Adquirente            ReplaceKey = "{{ADQUIRENTE}}"
	TabelionatoNum        ReplaceKey = "{{TABELIONATO_NUM}}"
	TabelionatoName       ReplaceKey = "{{TABELIONATO_NAME}}"
	TabelionatoCityState  ReplaceKey = "{{TABELIONATO_CITY_STATE}}"
	BookNum               ReplaceKey = "{{BOOK_NUM}}"
	BookPages             ReplaceKey = "{{BOOK_PAGES}}"
	EscrituraMadeDate     ReplaceKey = "{{BOOK_DATE}}"
	EscrituraValor        ReplaceKey = "{{ESCRITURA_VALOR}}"
	EscrituraValorExtenso ReplaceKey = "{{ESCRITURA_VALOR_EXTENSO}}"
	ItbiValor             ReplaceKey = "{{ITBI_VALOR}}"
	ItbiIncidenciaValor   ReplaceKey = "{{ITBI_INCIDENCIA_VALOR}}"
)

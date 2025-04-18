package minuta

// ReplaceKey -> é para substiuir elementos no minuta template
// Eu preciso relacionar A chave que eu irei substituir no template
// Com as propriedades necessário para achar o valor dessa chave

// Então podemos dizer que um Template Tem N chaves para serem substituídas
// Cada chave sabe dizer qual seu formato a ser substituído no template
// Cada chave sabe dizer qual é seu Start e End Key para encontrarmos o valor dessa chave

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

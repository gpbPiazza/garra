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
	TypeAtoRK       ReplaceKey = "{{TYPE_ATO}}"
	NumAtoRK        ReplaceKey = "{{NUM_ATO}}"
	TitleAtoRK      ReplaceKey = "{{TITLE_ATO}}"
	MatriculaRK     ReplaceKey = "{{MATRICULA}}"
	DataRegistroRK  ReplaceKey = "{{DATA_REGISTRO}}"
	ProtocoloRK     ReplaceKey = "{{PROTOCOLO}}"
	DataProtocoloRK ReplaceKey = "{{DATA_PROTOCOLO}}"

	// MINUTA KEYS
	TransmitenteRK ReplaceKey = "{{TRANSMITENTE}}"
	AdquirenteRK   ReplaceKey = "{{ADQUIRENTE}}"

	TabelionatoNameRK     ReplaceKey = "{{TABELIONATO_NAME}}"
	TabelionatoCityUFRK   ReplaceKey = "{{TABELIONATO_CITY_STATE}}"
	InitialBookPagesRK    ReplaceKey = "{{INITIAL_BOOK_PAGES}}"
	FinalBookPagesRK      ReplaceKey = "{{FINAL_BOOK_PAGES}}"
	BookNumRK             ReplaceKey = "{{BOOK_NUM}}"
	EscrituraMadeDateRK   ReplaceKey = "{{BOOK_DATE}}"
	EscrituraValorRK      ReplaceKey = "{{ESCRITURA_VALOR}}"
	ItbiValorRK           ReplaceKey = "{{ITBI_VALOR}}"
	ItbiIncidenciaValorRK ReplaceKey = "{{ITBI_INCIDENCIA_VALOR}}"
)

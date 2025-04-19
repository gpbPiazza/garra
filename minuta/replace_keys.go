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
	Transmitente ReplaceKey = "{{TRANSMITENTE}}" //TODO: REMOVE THIS KEYS
	Adquirente   ReplaceKey = "{{ADQUIRENTE}}"   //TODO: REMOVE THIS KEYS

	// TRANSMITENTE = Outorgante
	TransmitenteNome             ReplaceKey = "{{TRANSMITENTE_NOME}}"
	TransmitenteJob              ReplaceKey = "{{TRANSMITENTE_JOB}}"
	TransmitenteNacionalidade    ReplaceKey = "{{TRANSMITENTE_NACIONALIDADE}}"
	TransmitenteEstadoCivil      ReplaceKey = "{{TRANSMITENTE_ESTADO_CIVIL}}"
	TransmitenteCPF_CNPJ         ReplaceKey = "{{TRANSMITENTE_CPF_CNPJ}}"
	TransmitenteEnderecoRua      ReplaceKey = "{{TRANSMITENTE_ENDERECO_RUA}}"
	TransmitenteEnderecoN        ReplaceKey = "{{TRANSMITENTE_ENDERECO_NUMERO}}"
	TransmitenteEnderecoBairro   ReplaceKey = "{{TRANSMITENTE_ENDERECO_BAIRRO}}"
	TransmitenteEnderecoCidadeUF ReplaceKey = "{{TRANSMITENTE_ENDERECO_CIDADE_UF}}"

	// Aduirente = Outorgado
	AdquirenteNome             ReplaceKey = "{{ADQUIRENTE_NOME}}"
	AdquirenteJob              ReplaceKey = "{{ADQUIRENTE_JOB}}"
	AdquirenteNacionalidade    ReplaceKey = "{{ADQUIRENTE_NACIONALIDADE}}"
	AdquirenteEstadoCivil      ReplaceKey = "{{ADQUIRENTE_ESTADO_CIVIL}}"
	AdquirenteCPF_CNPJ         ReplaceKey = "{{ADQUIRENTE_CPF_CNPJ}}"
	AdquirenteEnderecoRua      ReplaceKey = "{{ADQUIRENTE_ENDERECO_RUA}}"
	AdquirenteEnderecoN        ReplaceKey = "{{ADQUIRENTE_ENDERECO_NUMERO}}"
	AdquirenteEnderecoBairro   ReplaceKey = "{{ADQUIRENTE_ENDERECO_BAIRRO}}"
	AdquirenteEnderecoCidadeUF ReplaceKey = "{{ADQUIRENTE_ENDERECO_CIDADE_UF}}"

	TabelionatoName      ReplaceKey = "{{TABELIONATO_NAME}}"
	TabelionatoCityState ReplaceKey = "{{TABELIONATO_CITY_STATE}}"
	BookNum              ReplaceKey = "{{BOOK_NUM}}"
	InitialBookPage      ReplaceKey = "{{INITIAL_BOOK_PAGE}}"
	FinalBookPage        ReplaceKey = "{{FINAL_BOOK_PAGE}}"
	EscrituraMadeDate    ReplaceKey = "{{BOOK_DATE}}"
	EscrituraValor       ReplaceKey = "{{ESCRITURA_VALOR}}"
	ItbiValor            ReplaceKey = "{{ITBI_VALOR}}"
	ItbiIncidenciaValor  ReplaceKey = "{{ITBI_INCIDENCIA_VALOR}}"
)

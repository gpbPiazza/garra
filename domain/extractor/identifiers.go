package extractor

type Identifier int

const (
	TypeAtoID Identifier = iota
	NumAtoID
	TitleAtoID
	MatriculaID
	DataRegistroID
	ProtocoloID
	DataProtocoloID

	OutorganteID
	OutorgadoID

	OutorganteName
	OutorganteJob
	OutorganteSex
	OutorganteNationality
	OutorganteEstadoCivil
	OutorganteDocNumCPF_CNPJ
	OutorganteDocType
	OutorganteEnderecoRua
	OutorganteEnderecoN
	OutorganteEnderecoBairro
	OutorganteEnderecoCidadeUF

	OutorgadoName
	OutorgadoJob
	OutorgadoSex
	OutorgadoNationality
	OutorgadoEstadoCivil
	OutorgadoDocNumCPF_CNPJ
	OutorgadoDocType
	OutorgadoEnderecoRua
	OutorgadoEnderecoN
	OutorgadoEnderecoBairro
	OutorgadoEnderecoCidadeUF

	TabelionatoNameID
	TabelionatoCityUFID
	InitialBookPagesID
	FinalBookPagesID
	BookNumID
	EscrituraMadeDateID
	EscrituraValorID
	ItbiValorID
	ItbiIncidenciaValorID

	NameID
	JobID
	SexID
	NationalityID
	MaritialStatusID
	DocNumCPF_CNPJID
	DocTypeID
	AddressStreetID
	AddressNID
	AddressNeighborhoodID
	AddressCityUFID

	InFavorToWhoID
	WhoDoesID
)

var IdentifiersNames = map[Identifier]string{
	TypeAtoID:                  "Tipo do Ato",
	NumAtoID:                   "Número do Ato",
	TitleAtoID:                 "Título do Ato",
	MatriculaID:                "Matrícula",
	DataRegistroID:             "Data de registro",
	ProtocoloID:                "Protocolo",
	DataProtocoloID:            "Data do protocolo",
	OutorganteID:               "Outorgante",
	OutorgadoID:                "Outorgado",
	OutorganteName:             "Outorgante nome",
	OutorganteJob:              "Outorgante trabalho",
	OutorganteNationality:      "Outorgante nacionalidade",
	OutorganteEstadoCivil:      "Outorgante estado Civil",
	OutorganteDocNumCPF_CNPJ:   "Outorgante CPF ou CNPJ",
	OutorganteEnderecoRua:      "Outorgante endereço rua",
	OutorganteEnderecoN:        "Outorgante número do endereço",
	OutorganteEnderecoBairro:   "Outorgante endereço bairro",
	OutorganteEnderecoCidadeUF: "Outorgante endereço cidade e UF",
	OutorgadoName:              "Outorgado nome",
	OutorgadoJob:               "Outorgado trabalho",
	OutorgadoNationality:       "Outorgado naciolidade",
	OutorgadoEstadoCivil:       "Outorgado estado civil",
	OutorgadoDocNumCPF_CNPJ:    "Outorgado CPF ou CNPJ",
	OutorgadoEnderecoRua:       "Outorgado endereço rua",
	OutorgadoEnderecoN:         "Outorgado número do endereço",
	OutorgadoEnderecoBairro:    "Outorgado endereço bairro",
	OutorgadoEnderecoCidadeUF:  "Outorgado endereço cidade e UF",
	TabelionatoNameID:          "Tabelionato nome",
	TabelionatoCityUFID:        "Tabelionato cidade e UF",
	InitialBookPagesID:         "Páginas iniciais do livro",
	FinalBookPagesID:           "Páginas finais do livro",
	BookNumID:                  "Número do livro",
	EscrituraMadeDateID:        "Data da escritura",
	EscrituraValorID:           "Valor da escrita",
	ItbiValorID:                "valor do ITBI",
	ItbiIncidenciaValorID:      "Valor da incidência do ITBI",
	//
	NameID:                "Nome",
	JobID:                 "Trabalho",
	SexID:                 "Sexo",
	NationalityID:         "Nacionalidade",
	MaritialStatusID:      "Estado Civil",
	DocNumCPF_CNPJID:      "Número do documento CPF_CNPJ",
	DocTypeID:             "Tipo do documento",
	AddressStreetID:       "Rua do endereço",
	AddressNID:            "Número do endereço",
	AddressNeighborhoodID: "Bairro do endereço",
	AddressCityUFID:       "CidadeUF do endereço",

	InFavorToWhoID: "Em favor de quem",
	WhoDoesID:      "Quem faz",
}

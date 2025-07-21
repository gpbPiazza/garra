package extractor

type Identifier int

const (
	TypeAto Identifier = iota
	NumAto
	TitleAto
	Matricula
	DataRegistro
	Protocolo
	DataProtocolo

	Outorgante
	Outorgado

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

	TabelionatoName
	TabelionatoCityUF
	InitialBookPages
	FinalBookPages
	BookNum
	EscrituraMadeDate
	EscrituraValor
	ItbiValor
	ItbiIncidenciaValor

	Name
	Job
	Sex
	Nationality
	MaritialStatus
	DocNumCPF_CNPJ
	DocType
	AddressStreet
	AddressN
	AddressNeighborhood
	AddressCityUF
)

var IdentifiersNames = map[Identifier]string{
	TypeAto:                    "Tipo do Ato",
	NumAto:                     "Número do Ato",
	TitleAto:                   "Título do Ato",
	Matricula:                  "Matrícula",
	DataRegistro:               "Data de registro",
	Protocolo:                  "Protocolo",
	DataProtocolo:              "Data do protocolo",
	Outorgante:                 "Outorgante",
	Outorgado:                  "Outorgado",
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
	TabelionatoName:            "Tabelionato nome",
	TabelionatoCityUF:          "Tabelionato cidade e UF",
	InitialBookPages:           "Páginas iniciais do livro",
	FinalBookPages:             "Páginas finais do livro",
	BookNum:                    "Número do livro",
	EscrituraMadeDate:          "Data da escritura",
	EscrituraValor:             "Valor da escrita",
	ItbiValor:                  "valor do ITBI",
	ItbiIncidenciaValor:        "Valor da incidência do ITBI",
	//
	Name:                "Nome",
	Job:                 "Trabalho",
	Sex:                 "Sexo",
	Nationality:         "Nacionalidade",
	MaritialStatus:      "Estado Civil",
	DocNumCPF_CNPJ:      "Número do documento CPF_CNPJ",
	DocType:             "Tipo do documento",
	AddressStreet:       "Rua do endereço",
	AddressN:            "Número do endereço",
	AddressNeighborhood: "Bairro do endereço",
	AddressCityUF:       "CidadeUF do endereço",
}

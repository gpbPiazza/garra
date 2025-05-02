package extractor

type ResultKey int

const (
	TypeAto ResultKey = iota
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
)

var resultKeyNames = map[ResultKey]string{
	TypeAto:                    "TypeAto",
	NumAto:                     "NumAto",
	TitleAto:                   "TitleAto",
	Matricula:                  "Matricula",
	DataRegistro:               "DataRegistro",
	Protocolo:                  "Protocolo",
	DataProtocolo:              "DataProtocolo",
	Outorgante:                 "Outorgante",
	Outorgado:                  "Outorgado",
	OutorganteName:             "OutorganteName",
	OutorganteJob:              "OutorganteJob",
	OutorganteNationality:      "OutorganteNationality",
	OutorganteEstadoCivil:      "OutorganteEstadoCivil",
	OutorganteDocNumCPF_CNPJ:   "OutorganteCPF_CNPJ",
	OutorganteEnderecoRua:      "OutorganteEnderecoRua",
	OutorganteEnderecoN:        "OutorganteEnderecoN",
	OutorganteEnderecoBairro:   "OutorganteEnderecoBairro",
	OutorganteEnderecoCidadeUF: "OutorganteEnderecoCidadeUF",
	OutorgadoName:              "OutorgadoName",
	OutorgadoJob:               "OutorgadoJob",
	OutorgadoNationality:       "OutorgadoNationality",
	OutorgadoEstadoCivil:       "OutorgadoEstadoCivil",
	OutorgadoDocNumCPF_CNPJ:    "OutorgadoCPF_CNPJ",
	OutorgadoEnderecoRua:       "OutorgadoEnderecoRua",
	OutorgadoEnderecoN:         "OutorgadoEnderecoN",
	OutorgadoEnderecoBairro:    "OutorgadoEnderecoBairro",
	OutorgadoEnderecoCidadeUF:  "OutorgadoEnderecoCidadeUF",
	TabelionatoName:            "TabelionatoName",
	TabelionatoCityUF:          "TabelionatoCityState",
	InitialBookPages:           "InitialBookPages",
	FinalBookPages:             "FinalBookPages",
	BookNum:                    "BookNum",
	EscrituraMadeDate:          "EscrituraMadeDate",
	EscrituraValor:             "EscrituraValor",
	ItbiValor:                  "ItbiValor",
	ItbiIncidenciaValor:        "ItbiIncidenciaValor",
}

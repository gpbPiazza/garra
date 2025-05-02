package minuta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinuta(t *testing.T) {
	t.Run("Compra e venda minuta good params", func(t *testing.T) {
		params := MinutaParams{
			Transmitente: PersonParams{
				Name:            "RUZZU CONSTRUTORA E INCORPORADORA LTDA.",
				IsOverqualified: true,
			},
			Adquirente: PersonParams{
				Name:            "SIDNEI ANTÔNIO GATTIS",
				Nationality:     "Brasil",
				MaritalStatus:   "solteiro",
				DocNum_CPF_CNPJ: "03756166910",
				DocType:         "CPF",
				Address: AddressParams{
					Rua:          "Rua Azambuja",
					Num:          "541",
					Neighborhood: "Azambuja",
					CityUF:       "Brusque/SC",
				},
			},
			TitleAto:            "Compra e Venda",
			TabelionatoName:     "1º Tabelionato de Notas e Protesto",
			TabelionatoCityUF:   "Brusque/SC",
			BookNum:             "965",
			InitialBookPages:    "121",
			FinalBookPages:      "123",
			EscrituraMadeDate:   "26/03/2025",
			EscrituraValor:      "120.000,00 (cento e vinte mil reais)",
			ItbiValor:           "2.900,00",
			ItbiIncidenciaValor: "145.000,00",
		}

		expected := `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>RUZZU CONSTRUTORA E INCORPORADORA LTDA., supraqualificada.</span>.
<br/>
<u>ADQUIRENTE(S)</u>:><span>SIDNEI ANTÔNIO GATTIS, brasileiro, solteiro, CPF nº 037.561.669-10, residente e domiciliado na Rua Azambuja, nº 541, Bairro Azambuja, Brusque/SC.</span>.
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e Protesto de Brusque/SC, Livro 965, Folhas 121/123V, em 26/03/2025. 
<br/><u>VALOR</u>: R$ 120.000,00 (cento e vinte mil reais).
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ 2.900,00, com incidência sobre R$ 145.000,00, devidamente quitado. Nos termos do artigo 320 do CNCGFE/SC, o imóvel da presente matrícula, teve como valor atribuído de mercado, no
''quantum'' de R$ XXXXX. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

		got := Minuta(params)

		assert.Equal(t, expected, got)
	})
}

func TestMinutaPerson(t *testing.T) {
	t.Run("Valid person with CPF", func(t *testing.T) {
		person := PersonParams{
			Name:            "SIDNEI ANTÔNIO GATTIS",
			Nationality:     "Brasil",
			MaritalStatus:   "solteiro",
			DocNum_CPF_CNPJ: "03756166910",
			DocType:         "CPF",
			Address: AddressParams{
				Rua:          "Rua Azambuja",
				Num:          "541",
				Neighborhood: "Azambuja",
				CityUF:       "Brusque/SC",
			},
		}

		expected := "SIDNEI ANTÔNIO GATTIS, brasileiro, solteiro, CPF nº 037.561.669-10, residente e domiciliado na Rua Azambuja, nº 541, Bairro Azambuja, Brusque/SC."

		got, err := minutaPerson(person)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("valid person with CNPJ", func(t *testing.T) {
		person := PersonParams{
			Name:            "Some name",
			Nationality:     "Brasil",
			MaritalStatus:   "solteiro",
			DocNum_CPF_CNPJ: "12345678000195",
			DocType:         "CNPJ",
			Address: AddressParams{
				Rua:          "Rua Azambuja",
				Num:          "541",
				Neighborhood: "Azambuja",
				CityUF:       "Brusque/SC",
			},
		}

		// TODO: ajust this for CNPJ persons
		expected := "Some name, brasileiro, solteiro, CPF nº 12.345.678/0001-95, residente e domiciliado na Rua Azambuja, nº 541, Bairro Azambuja, Brusque/SC."

		got, err := minutaPerson(person)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("return supraqualificada when person IsOverqualified", func(t *testing.T) {
		person := PersonParams{
			Name:            "RUZZU CONSTRUTORA E INCORPORADORA LTDA.",
			DocNum_CPF_CNPJ: "12345678000195",
			DocType:         "CNPJ",
			IsOverqualified: true,
		}

		expected := "RUZZU CONSTRUTORA E INCORPORADORA LTDA., supraqualificada."

		got, err := minutaPerson(person)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Invalid CPF format only 10 digits", func(t *testing.T) {
		person := PersonParams{
			Name:            "JOÃO DA SILVA",
			Nationality:     "Brasil",
			MaritalStatus:   "solteiro",
			DocNum_CPF_CNPJ: "0375616691",
			DocType:         "CPF",
			Address: AddressParams{
				Rua:          "Rua Azambuja",
				Num:          "541",
				Neighborhood: "Bairro Azambuja",
				CityUF:       "Brusque/SC",
			},
		}

		got, err := minutaPerson(person)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "malformed CPF value len diff than 11")
		assert.Empty(t, got)
	})

	t.Run("Invalid CNPJ format only 13 digits", func(t *testing.T) {
		person := PersonParams{
			Name:            "EMPRESA LTDA.",
			DocNum_CPF_CNPJ: "1234567890123",
			DocType:         "CNPJ",
		}

		got, err := minutaPerson(person)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "malformed CNPJ value len diff than 14")
		assert.Empty(t, got)
	})

	t.Run("UNKNOWN docType", func(t *testing.T) {
		person := PersonParams{
			Name:            "EMPRESA LTDA.",
			DocNum_CPF_CNPJ: "1234567890123",
			DocType:         "UNKNOWN",
		}

		got, err := minutaPerson(person)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "docType not mapped - type UNKNOWN")
		assert.Empty(t, got)
	})

	t.Run("Invalid nationality", func(t *testing.T) {
		person := PersonParams{
			Name:            "JOÃO DA SILVA",
			Nationality:     "Argentina",
			MaritalStatus:   "solteiro",
			DocNum_CPF_CNPJ: "03756166910",
			DocType:         "CPF",
			Address: AddressParams{
				Rua:          "Rua Azambuja",
				Num:          "541",
				Neighborhood: "Bairro Azambuja",
				CityUF:       "Brusque/SC",
			},
		}

		got, err := minutaPerson(person)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "nationality not mapped - got Argentina")
		assert.Empty(t, got)
	})

	t.Run("Invalid cityUF format - missing slash", func(t *testing.T) {
		person := PersonParams{
			Name:            "JOÃO DA SILVA",
			Nationality:     "Brasil",
			MaritalStatus:   "solteiro",
			DocNum_CPF_CNPJ: "03756166910",
			DocType:         "CPF",
			Address: AddressParams{
				Rua:          "Rua Azambuja",
				Num:          "541",
				Neighborhood: "Bairro Azambuja",
				CityUF:       "Brusque SC",
			},
		}

		got, err := minutaPerson(person)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "cityUF can not be splited by /")
		assert.Empty(t, got)
	})

	t.Run("Invalid cityUF format - empty string", func(t *testing.T) {
		person := PersonParams{
			Name:            "JOÃO DA SILVA",
			Nationality:     "Brasil",
			MaritalStatus:   "solteiro",
			DocNum_CPF_CNPJ: "03756166910",
			DocType:         "CPF",
			Address: AddressParams{
				Rua:          "Rua Azambuja",
				Num:          "541",
				Neighborhood: "Bairro Azambuja",
				CityUF:       "",
			},
		}

		got, err := minutaPerson(person)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "cityUF can not be splited by /")
		assert.Empty(t, got)
	})

	t.Run("Valid cityUF input with extra spaces", func(t *testing.T) {
		person := PersonParams{
			Name:            "JOÃO DA SILVA",
			Nationality:     "BRASIL",
			MaritalStatus:   "solteiro",
			DocNum_CPF_CNPJ: "03756166910",
			DocType:         "CPF",
			Address: AddressParams{
				Rua:          "Rua Azambuja",
				Num:          "541",
				Neighborhood: "Azambuja",
				CityUF:       "  Brusque / SC  ",
			},
		}

		expected := "JOÃO DA SILVA, brasileiro, solteiro, CPF nº 037.561.669-10, residente e domiciliado na Rua Azambuja, nº 541, Bairro Azambuja, Brusque/SC."

		got, err := minutaPerson(person)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

package minuta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

		expected := "Some name., CNPJ nº 12.345.678/0001-95, com sede na rua Rua Azambuja, nº 541, Bairro Azambuja, Brusque/SC."

		got, err := minutaPerson(person)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("return supraqualificada when person IsOverqualified", func(t *testing.T) {
		person := PersonParams{
			Name:            "RUZZU CONSTRUTORA E INCORPORADORA LTDA",
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

func TestFormatValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Missing closing parenthesis with date",
			input:    "R$ 200.000,00 (duzentos mil24/04/2025.",
			expected: "R$ 200.000,00 (duzentos mil reais)",
			wantErr:  false,
		},
		{
			name:     "Typo in 'e dez'",
			input:    "R$210.000,00 (duzentos edez mil reais)",
			expected: "R$ 210.000,00 (duzentos e dez mil reais)",
			wantErr:  false,
		},
		{
			name:     "Correct format already",
			input:    "R$210.000,00 (duzentos e dez mil reais)",
			expected: "R$ 210.000,00 (duzentos e dez mil reais)",
			wantErr:  false,
		},
		{
			name:     "Different amount but correct format",
			input:    "R$120.000,00 (cento e vinte mil reais)",
			expected: "R$ 120.000,00 (cento e vinte mil reais)",
			wantErr:  false,
		},
		{
			name:     "Empty input",
			input:    "",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Missing parenthesis",
			input:    "R$ 200.000,00 duzentos mil reais",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatValue(tt.input)

			// Check error cases
			if (err != nil) != tt.wantErr {
				t.Errorf("formatValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check results when no error expected
			if !tt.wantErr && got != tt.expected {
				t.Errorf("formatValue() got = %q, want %q", got, tt.expected)
			}
		})
	}
}

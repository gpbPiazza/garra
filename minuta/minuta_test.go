package minuta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinuta(t *testing.T) {
	t.Run("Compra e venda minuta good params", func(t *testing.T) {
		params := MinutaParams{
			Transmitente:          "RUZZU CONSTRUTORA E INCORPORADORA LTDA., supraqualificada.",
			Adquirente:            "SIDNEI ANTÔNIO GATTIS, brasileiro, solteiro, CPF nº 037.561.669-10, residente e domiciliado na Rua Azambuja, nº 541, Bairro Azambuja, Brusque/SC.",
			TitleAto:              "Compra e Venda",
			TabelionatoNum:        "1º",
			TabelionatoName:       "Tabelionato de Notas e Protesto",
			TabelionatoCityState:  "Brusque/SC",
			BookNum:               "965",
			BookPages:             "121/123V",
			EscrituraMadeDate:     "26/03/2025",
			EscrituraValor:        "120.000,00",
			EscrituraValorExtenso: "(cento e vinte mil reais)",
			ItbiValor:             "2.900,00",
			ItbiIncidenciaValor:   "145.000,00",
		}
		//Adiquirente nome
		//Adiquirente nacionalidade
		//Adiquirente estado civil
		//Adiquirente CPF/CPNJ
		//Adiquirente Endereco Rua
		// Adiquirente Endereco N
		// Adiquirente Endereco bairro
		// Adiquirente Endereco Cidade
		// Adiquirente Endereco Estado

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
''quantum'' de R$ XXXXX. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com
as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

		got := Minuta(params)

		assert.Equal(t, expected, got)
	})
}

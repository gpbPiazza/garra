package minuta

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("Generate good Minuta from Good Document", func(t *testing.T) {
		generatorApp := NewGeneratorApp()

		expected := `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>CONSTRUARTE CONSTRUTORA E INCORPORADORA LTDA, Brasil, Casado (a), CPF nº 2923597700017014, residente e domiciliado na Rua Luiz Eccel, nº 1, Paquetá, Brusque / SC.</span>.
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

		doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_1_to_1_buy_and_sell.txt")
		require.NoError(t, err)

		got, err := generatorApp.Generate(string(doc))
		require.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}
